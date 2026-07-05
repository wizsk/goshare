package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

// ipList lets -allow-ip be passed multiple times, e.g.
//
//	goshare --allow-ip 192.168.1.5 --allow-ip 192.168.1.10
//
// If it's never set (len == 0), the middleware allows every IP through,
// so behavior is unchanged for anyone not using the flag.
var allowedIPSlice []string

func populateIpListSlice(ipListString string) {

	if ipListString == "" {
		return
	}
	if strings.ContainsAny(ipListString, ",") {

		ipaddrs := strings.Split(ipListString, ",")

		for i, v := range ipaddrs {
			ip := strings.TrimSpace(v)
			if net.ParseIP(ip) == nil {
				log.Fatalf("--allow-ip: %q is not a valid ip address", ip)
			}
			allowedIPSlice = append(allowedIPSlice, ip)
			if i+1 == len(ipaddrs) {
				allowedIPSlice = append(allowedIPSlice, "127.0.0.1")
				return
			}
		}
	}
	allowedIPSlice = append(allowedIPSlice, "127.0.0.1")
	allowedIPSlice = append(allowedIPSlice, strings.TrimSpace(ipListString))
}

// remoteIP strips the port off r.RemoteAddr, returning just the IP.
func remoteIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// RemoteAddr didn't have a port for some reason, just use it as-is
		return r.RemoteAddr
	}
	return host
}

// ipAllowed reports whether addr is allowed to make requests.
// With no -allow-ip flags set, everything is allowed (default, unrestricted behavior).
func ipAllowed(addr string) bool {
	if len(allowedIPSlice) == 0 {
		return true
	}
	for _, ip := range allowedIPSlice {
		if ip == addr {
			return true
		}
	}
	return false
}

// restrictIP wraps a handler so it 403s any request whose IP isn't in
// the -allow-ip list (when that list is non-empty).
func restrictIP(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := remoteIP(r)
		if !ipAllowed(addr) {
			http.Error(w, "403 forbidden: your ip is not allowed to access this server", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

///
