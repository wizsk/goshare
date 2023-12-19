package main

import (
	"net/http"

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

// funcs for authenticated usrs
func (s *server) authBrowse(w http.ResponseWriter, r *http.Request)  { auth(w, r, s.browse) }
func (s *server) authDownZip(w http.ResponseWriter, r *http.Request) { auth(w, r, s.downZip) }
func (s *server) authMkdir(w http.ResponseWriter, r *http.Request)   { auth(w, r, s.mkdir) }
func (s *server) authUpload(w http.ResponseWriter, r *http.Request)  { auth(w, r, s.upload) }
func (s *server) authZip(w http.ResponseWriter, r *http.Request)     { auth(w, r, s.zip) }

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
