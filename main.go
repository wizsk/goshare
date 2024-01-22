package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const debug = !false

var (
	rootDir, port, password string

	showStat bool
)

func flagParse() {
	flag.StringVar(&rootDir, "d", ".", "the directory for sharing")
	flag.StringVar(&port, "port", "8001", "port number")
	flag.StringVar(&password, "p", "", "password")
	flag.BoolVar(&showStat, "s", false, "don't show request information. aka silent")
	flag.Parse()

	if port == "" {
		log.Fatal("prot number can't be empty")
	}

	if rootDir == "" {
		log.Fatal("directory name can't be empty")
	}
}

func localIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "localhost"
}

func main() {
	flagParse()

	sv := newServer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	// don't chage the /browse/ ok it will break suff
	http.HandleFunc("/auth", sv.aunth)
	http.HandleFunc("/browse/", sv.authBrowse)
	http.HandleFunc("/zip", sv.authZip)
	http.HandleFunc("/downzip/", sv.authDownZip)
	http.HandleFunc("/upload", sv.authUpload)
	http.HandleFunc("/mkdir", sv.authMkdir)

	fmt.Printf("serving at http://%s:%s\n", localIp(), port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		os.Exit(1)
	}
}
