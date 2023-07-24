package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wizsk/goshare/auth"
)

const version = "v1.1"

var dir = flag.String("d", ".", "direcotry name")
var port = flag.String("port", "8001", "port number")
var pass = flag.String("p", "", "password")
var verstionFlag = flag.Bool("v", false, "prints current version")

func main() {
	flag.Parse()
	if *verstionFlag {
		fmt.Printf("goshare current version: %s\n", version)
		os.Exit(0)
	}

	fileSeverInit(*dir)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "sorry someting went wrong", http.StatusBadRequest)
			return
		}

		// useing ?query=string to avoid making atoher handeler
		// http://example.com/file?cli=password
		if cli := r.FormValue("cli"); cli != "" {
			if cli == *pass || *pass == "" {
				cliUi(w, r, cli)
			} else {
				http.Error(w, "please provide as such http://example.com/file?cli=password", http.StatusBadRequest)
			}
			return
		}

		if res := r.FormValue("res"); res != "" {
			serveResource(w, res)
			return
		}

		if *pass != "" {
			if r.Method == http.MethodPost {
				if r.FormValue("password") != *pass {
					serveResource(w, "form")
					return
				}
				auth.WriteCookie(w)
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
				return
			}

			if err := auth.ReadCookie(r, auth.CookieName); err != nil {
				serveResource(w, "form")
				return
			}
		}

		ServeWebUi(w, r)
	})

	if *pass == "" {
		fmt.Printf("serving %q at http://localhost:%s\n", *dir, *port)
	} else {
		fmt.Printf("serving %q at http://localhost:%s\npassword: %s\n", *dir, *port, *pass)
	}

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
