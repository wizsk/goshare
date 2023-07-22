package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wizsk/goshare/auth"
)

var dir = flag.String("d", "", "direcotry name")
var port = flag.String("port", "8001", "port number")
var pass = flag.String("p", "", "port number")

const randomPath = "/dQw4w9WgXcQ"

func main() {
	flag.Parse()
	if *dir == "" {
		fmt.Println("error: directory name not provided")
		fmt.Print(usages)
		os.Exit(1)
	}
	FileSeverInit(*dir)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *pass == "" {
			ServeFiles(w, r)
			return
		}
		if err := auth.ReadCookie(r, auth.CookieName); err != nil {
			w.Write(form)
			return
		}
		ServeFiles(w, r)
	})

	http.HandleFunc(randomPath + "/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path[len(randomPath):] {
		case "/favicon.svg":
			w.Write(favicon)
		case "/auth":
			authHelper(w, r)
		case "/css":
			http.ServeFile(w,r,"tailwind/output.css")
		}
	})

	fmt.Println("serving at http://localhost:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
