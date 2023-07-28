package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/wizsk/goshare/auth"
	"github.com/wizsk/goshare/compress"
)

const version = "v2.0"
const debug = false

var ZIP_DIR string

var dir = flag.String("d", ".", "direcotry name")
var port = flag.String("port", "8001", "port number")
var pass = flag.String("p", "", "password")
var verstionFlag = flag.Bool("v", false, "prints current version")

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

	go func() {
		<-sighalChannel
		fmt.Println("terminating server and cleaning temp files")
		err := os.RemoveAll(ZIP_DIR)
		if err != nil {
			log.Println(err)
		}
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
	if cli := r.FormValue("cli"); cli != "" {
		if cli == *pass || *pass == "" {
			ServeWebUi(w, r)
		} else {
			printStat(r, LOGIN_ATTEMT)
			http.Error(w, "please provide as such http://example.com/file?cli=password", http.StatusBadRequest)
		}
		return
	}

	// if password was set then authorize
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
		printStat(r, ZIP_REQUEST)
		serveZipFile(w, r, zip)
		return
	}

	printStat(r, NORMAL_REQUEST)
	ServeWebUi(w, r)
}
