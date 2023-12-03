package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed frontend/src/*
var templateFiles embed.FS

type server struct {
	root, tmp, zipSavePath string
	// tmpl      *template.Template
}

type svData struct {
	Dir  []Item
	Cd   string // current direcoty
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

	currentDirName := ""
	if stat, err := os.Stat(fileName); err != nil {
		log.Println(err)
		return
	} else if !stat.IsDir() {
		http.ServeFile(w, r, fileName)
		return
	} else {
		if r.URL.Path == "/browse/" {
			currentDirName = "/"
		} else {
			currentDirName = stat.Name()
		}
	}

	indexPage := template.New("fo").Funcs(template.FuncMap{
		"pathJoin": filepath.Join,
	})

	var err error
	if debug {
		indexPage, err = indexPage.ParseGlob("frontend/src/*")
	} else {
		indexPage, err = indexPage.ParseFS(templateFiles, "frontend/src/*")
	}

	if err != nil {
		log.Panicln(err)
	}

	svd := svData{Od: r.URL.Path, Cd: currentDirName}

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
