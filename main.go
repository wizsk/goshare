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

func main() {
	flagParse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/browse/", browse)

	fmt.Printf("serving at http://%s:%s\n", localIp(), port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		os.Exit(1)
	}
}

func browse(w http.ResponseWriter, r *http.Request) {
	// example.com/fo/bar/bazz -> ["/fo/", "/fo/bar", "/fo/bar/bazz"]
	var raw []string
	for _, itm := range strings.Split(r.URL.EscapedPath(), "/") {
		if len(itm) == 0 {
			continue
		}

		if len(raw) == 0 {
			raw = append(raw, "/"+itm)
		} else {
			raw = append(raw, raw[len(raw)-1]+"/"+itm)
		}
	}

	fmt.Println(raw)
	// fileName := filepath.Join(rootDir, strings.TrimPrefix(r.URL.Path, "/browse"))
	// if stat, err := os.Stat(file); err != nil {
	// 	log.Println(err)
	// 	return
	// } else if !stat.IsDir() {
	// 	http.ServeFile(w, r, file)
	// 	return
	// }
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
