package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type server struct {
	root, tmp, zipSavePath string

	// tmpl      *template.Template
}

type svData struct {
	Dir  []Item
	Od   string // working directory
	Umap []Umap
}

type Umap struct {
	Name, Url string
}

func (s *server) browse(w http.ResponseWriter, r *http.Request) {
	// example.com/fo/bar/bazz -> ["/fo/", "/fo/bar", "/fo/bar/bazz"]
	fileName := filepath.Join(s.root, strings.TrimPrefix(r.URL.Path, "/browse"))

	// fmt.Println("filename:", fileName)

	if stat, err := os.Stat(fileName); err != nil {
		log.Println(err)
		return
	} else if !stat.IsDir() {
		http.ServeFile(w, r, fileName)
		return
	}

	indexPage := template.New("fo").Funcs(template.FuncMap{
		// "pathJoin": func(s, y string) string {
		// 	return filepath.Join(s, y)
		// },
		"pathJoin": filepath.Join,
	})

	indexPage, err := indexPage.ParseGlob("frontend/src/*.html")
	if err != nil {
		log.Fatal(err)
	}
	svd := svData{Od: r.URL.Path}

	svd.Dir, err = readDir(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	for _, itm := range strings.Split(r.URL.EscapedPath(), "/") {
		if len(itm) == 0 {
			continue
		}

		if len(svd.Umap) == 0 {
			svd.Umap = append(svd.Umap, Umap{itm, "/" + itm})
		} else {
			svd.Umap = append(svd.Umap, Umap{itm, svd.Umap[len(svd.Umap)-1].Url + "/" + itm})
		}

	}

	err = indexPage.ExecuteTemplate(w, "index.html", &svd)
	if err != nil {
		log.Println(err)
		return
	}
}
