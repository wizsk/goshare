package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
)

var rootDir, port string

func flagParse() {
	flag.StringVar(&rootDir, "d", ".", "the directory for sharing")
	flag.StringVar(&port, "p", "8001", "port number")
	flag.Parse()
}

func main() {
	flagParse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse", http.StatusMovedPermanently)
	})

	http.HandleFunc("/browse", browse)

	fmt.Printf("serving at http://%s:%s\n", localIp(), port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		os.Exit(1)
	}

}

func browse(w http.ResponseWriter, r *http.Request) {
	buff := new(bytes.Buffer)

	http.Redirect(rootDir)

	w.Write([]byte("hi"))
	// fo
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
