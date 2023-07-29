package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/wizsk/goshare/auth"
	"github.com/wizsk/goshare/compress"
)

const version = "v2.1"
const debug = false

var ZIP_DIR string

var dir = flag.String("d", ".", "direcotry name")
var port = flag.String("port", "8001", "port number")
var pass = flag.String("p", "", "password")
var verstionFlag = flag.Bool("v", false, "prints current version")

var showStat = flag.Bool("s", false, "silence print informating about requests")

// var showStat bool

// for ptintstat
const (
	NORMAL_REQUEST = "browseing"
	ZIP_REQUEST    = "Zipping file"
	LOGIN_ATTEMT   = "login attemt"
	LOGIN_SUCCESS  = "login success"
	FILE_DOWN      = "downloading"
)

func main() {
	flag.Parse()
	// ican't find a bette way
	*showStat = !*showStat
	if *verstionFlag {
		fmt.Printf("goshare current version: %s\n", version)
		os.Exit(0)
	}

	fileSeverInit(*dir)

	sighalChannel := make(chan os.Signal, 1)
	signal.Notify(sighalChannel, os.Interrupt, syscall.SIGTERM)

	var err error
	ZIP_DIR, err = os.MkdirTemp(os.TempDir(), "goshare-zip-")
	if err != nil {
		log.Fatal(err)
	}
	compress.ZIP_PATH = ZIP_DIR

	// exiting gracefully
	go func() {
		<-sighalChannel
		fmt.Println("\nterminating server and cleaning temp files")
		err := os.RemoveAll(ZIP_DIR)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("waiting for backgorund prcoesses to stop")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	http.HandleFunc("/", mainHandeler)

	if *pass == "" {
		fmt.Printf("serving %q at http://localhost:%s\n\n", *dir, *port)
	} else {
		fmt.Printf("serving %q at http://localhost:%s\npassword: %s\n\n", *dir, *port, *pass)
	}
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

func mainHandeler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "sorry someting went wrong", http.StatusBadRequest)
		return
	}

	// useing ?query=string to avoid making more handelers
	if res := r.FormValue("res"); res != "" {
		serveResource(w, res)
		return
	}

	// http://localhost/file?cli=password for using or downloading from the cli

	// if password was set then authorize
	if *pass != "" {
		if cli := r.FormValue("cli"); cli != "" {
			if cli != *pass {
				if *showStat {
					printStat(r, LOGIN_ATTEMT)
				}
				http.Error(w, "please provide as such http://example.com/file?cli=password", http.StatusBadRequest)
				return
			}
		} else {
			redirectURL, _ := url.QueryUnescape(r.FormValue("redirect"))
			if redirectURL == "" {
				redirectURL = "/"
			}
			if r.Method == http.MethodPost {
				if r.FormValue("password") != *pass {
					serveForm(w, r)
					if *showStat {
						printStat(r, LOGIN_ATTEMT)
					}
					return
				}
				if *showStat {
					printStat(r, LOGIN_SUCCESS)
				}
				auth.WriteCookie(w)
				http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
				return
			}

			if err := auth.ReadCookie(r, auth.CookieName); err != nil {
				serveForm(w, r)
				return
			}
		}
	}

	if zip := r.FormValue("zip"); zip != "" {
		// printStat(r, ZIP_REQUEST)
		serveZipFile(w, r, zip)
		return
	}

	if *showStat {
		printStat(r, NORMAL_REQUEST)
	}
	if err := ServeWebUi(w, r); err != nil {
		if errors.Is(err, ErrNotDirectory) {
			fileUri := root + filepath.Clean(r.URL.Path)
			http.ServeFile(w, r, fileUri)
		} else if errors.Is(err, ErrFileNotFound) {
			http.Error(w, fmt.Sprintf("%q not found", r.URL.Path), http.StatusInternalServerError)
		} else {
			http.Error(w, "somethign went wrong", http.StatusInternalServerError)

		}
		return
	}
}
