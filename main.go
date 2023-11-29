package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var rootDir, port string

func flagParse() {
	flag.StringVar(&rootDir, "d", ".", "the directory for sharing")
	flag.StringVar(&port, "p", "8001", "port number")
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

	sv := server{rootDir, os.TempDir()}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/browse/", sv.browse)

	http.HandleFunc("/zip/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		fmt.Println()
		fmt.Println(r.Form["files"])
		for _, v := range r.Form["files"] {
			fmt.Println("\t", v, strings.Contains(v, ".."))
		}

		fmt.Println(r.FormValue("files"))
		fmt.Println()
	})

	fmt.Printf("serving at http://%s:%s\n", localIp(), port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		os.Exit(1)
	}
}
