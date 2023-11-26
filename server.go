package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type server struct {
	root, tmp string
	// tmpl      *template.Template
}

type svData struct {
	Dir []Item
	Od  string // working directory
}

func (s *server) browse(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("filename:", fileName)

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

	svd := svData{Od: r.URL.Path}

	svd.Dir, err = readDir(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	err = indexPage.Execute(w, &svd)
	if err != nil {
		log.Println(err)
		return
	}
}
