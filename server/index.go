package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (s *Server) Browse(w http.ResponseWriter, r *http.Request) {
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

	fileName := filepath.Join(s.root, strings.TrimPrefix(r.URL.Path, "/browse"))
	if stat, err := os.Stat(fileName); err != nil {
		log.Println(err)
		return
	} else if !stat.IsDir() {
		http.ServeFile(w, r, fileName)
		return
	}

	indexPage, err := template.ParseGlob("frontend/*.html")
	if err != nil {
		log.Fatal(err)
	}
	err = indexPage.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
