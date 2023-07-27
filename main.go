package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/wizsk/goshare/auth"
)

const version = "v2.0"
const styling = false

var dir = flag.String("d", ".", "direcotry name")
var port = flag.String("port", "8001", "port number")
var pass = flag.String("p", "", "password")
var verstionFlag = flag.Bool("v", false, "prints current version")

// for ptintstat
const (
	NORMAL_REQUEST = "browseing"
	LOGIN_ATTEMT   = "login attemt"
	LOGIN_SUCCESS  = "login success"
	FILE_DOWN      = "downloading"
)

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

		// titleasc, namedesc, sizeasc, sizedesc
		// useing ?query=string to avoid making atoher handeler
		// http://example.com/file?cli=password
		if cli := r.FormValue("cli"); cli != "" {
			if cli == *pass || *pass == "" {
				cliUi(w, r, cli)
			} else {
				printStat(r, LOGIN_ATTEMT)
				http.Error(w, "please provide as such http://example.com/file?cli=password", http.StatusBadRequest)
			}
			return
		}

		if res := r.FormValue("res"); res != "" {
			serveResource(w, res)
			return
		}

		if *pass != "" {
			redirectURL, _ := url.QueryUnescape(r.FormValue("redirect"))
			if redirectURL == "" {
				redirectURL = "/"
			}
			if r.Method == http.MethodPost {
				if r.FormValue("password") != *pass {
					serveForm(w, r)
					printStat(r, LOGIN_ATTEMT)
					return
				}
				printStat(r, LOGIN_SUCCESS)
				auth.WriteCookie(w)
				http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
				return
			}

			if err := auth.ReadCookie(r, auth.CookieName); err != nil {
				serveForm(w, r)
				return
			}
		}

		if zip := r.FormValue("zip"); zip != "" {
			serveZipFile(w, r, zip)
			return
		}

		printStat(r, NORMAL_REQUEST)
		ServeWebUi(w, r)
	})

	if *pass == "" {
		fmt.Printf("serving %q at http://localhost:%s\n\n", *dir, *port)
	} else {
		fmt.Printf("serving %q at http://localhost:%s\npassword: %s\n\n", *dir, *port, *pass)
	}

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
