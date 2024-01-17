package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wizsk/goshare/cookie"
)

func auth(w http.ResponseWriter, r *http.Request, handler func(w http.ResponseWriter, r *http.Request)) {
	if password != "" {
		if err := cookie.ReadCookie(r, cookie.CookieName); err != nil {
			http.Redirect(w, r, "/auth", http.StatusMovedPermanently)
			return
		}
	}
	handler(w, r)
}

// depends on a global var const
func (s *server) printStat(r *http.Request) {
	if s.showStat {
		rAddr := r.RemoteAddr
		if idx := strings.LastIndexByte(r.RemoteAddr, ':'); idx > 0 {
			rAddr = r.RemoteAddr[0:idx]
		}
		fmt.Printf("[REQ] %s | %15s | %6s | %q\n",
			time.Now().Format("2006/01/02 - 03:04:05 PM"),
			rAddr,
			r.Method,
			r.URL.Path,
		)
	}
}

func (s *server) authBrowse(w http.ResponseWriter, r *http.Request) {
	s.printStat(r)
	auth(w, r, s.browse)
}

func (s *server) authDownZip(w http.ResponseWriter, r *http.Request) {
	s.printStat(r)
	auth(w, r, s.downZip)
}

func (s *server) authMkdir(w http.ResponseWriter, r *http.Request) {
	s.printStat(r)
	auth(w, r, s.mkdir)
}

func (s *server) authUpload(w http.ResponseWriter, r *http.Request) {
	s.printStat(r)
	auth(w, r, s.upload)
}

func (s *server) authZip(w http.ResponseWriter, r *http.Request) {
	s.printStat(r)
	auth(w, r, s.zip)
}

// for "/auth" route
func (s *server) aunth(w http.ResponseWriter, r *http.Request) {
	if password == "" {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
		return
	}

	if r.Method != http.MethodPost {
		//	indexPage := template.New("fo").Funcs(template.FuncMap{ "pathJoin": filepath.Join, })
		w.Write([]byte(`<!DOCTYPE html> <html lang="en"><head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <title>Password Form</title> </head> <body> <form action="/auth" method="post"> <label for="password">Password:</label> <input type="password" id="password" name="password" required> <br> <input type="submit" value="Submit"> </form> </body> </html> `))
		return
	}

	pass := r.FormValue("password")
	if pass != password {
		http.Error(w, "password don't match", http.StatusBadRequest)
		return
	}

	if err := cookie.WriteCookie(w); err != nil {
		http.Error(w, "password don't match", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
}
