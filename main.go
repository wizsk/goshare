package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const debug = !false

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

	var zipD string
	var err error
	// var zipD = filepath.Join(os.TempDir(), "goshra_zip")

	if debug {
		fmt.Println("running in debug mode")
		zipD = filepath.Join(os.TempDir(), "goshare_zip")
		err = os.Mkdir(zipD, 0700)
		if os.IsExist(err) {
			err = nil
		}
	} else {
		zipD, err = os.MkdirTemp(os.TempDir(), "goshare_zip_")
	}
	if err != nil {
		log.Fatal(err)
	}

	if !debug {
		defer func() {
			os.RemoveAll(zipD)
		}()
	}

	sv := server{rootDir, os.TempDir(), zipD}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	// don't chage the /browse/ ok it will break suff
	http.HandleFunc("/browse/", sv.browse)
	http.HandleFunc("/zip", sv.zip)
	http.HandleFunc("/downzip/", sv.downZip)
	http.HandleFunc("/upload", sv.upload)
	http.HandleFunc("/mkdir", sv.mkdir)

	fmt.Printf("serving at http://%s:%s\n", localIp(), port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		os.Exit(1)
	}
}
