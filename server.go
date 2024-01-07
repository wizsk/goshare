package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

//go:embed frontend/src/*
var templateFiles embed.FS

type server struct {
	root, tmp, zipSavePath string
	showStat               bool
	// zipped                 map[string]string
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
		"timeFmt": func(t time.Time) string {
			return t.Format("01/02/2006 03:04 PM")
		},
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

	// sort.Slice(svd.Dir, func(i, j int) bool {
	// 	return svd.Dir[i].Name < svd.Dir[j].Name
	// })
	//
	// sort.Slice(svd.Dir, func(i, j int) bool {
	// 	return svd.Dir[i].IsDir || svd.Dir[j].IsDir
	// })

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

var validFilenameRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

func (s *server) mkdir(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	cwd := r.FormValue("cwd")
	if cwd == "" {
		http.Error(w, "cwd not provided", http.StatusBadRequest)
		return
	}

	parent := filepath.Join(s.root, strings.TrimPrefix(cwd, "/browse"))

	if pStat, err := os.Stat(parent); err != nil {
		http.Error(w, "could not resolve parent direcoty", http.StatusBadRequest)
		return
	} else if !pStat.IsDir() {
		http.Error(w, "paren is not a drectory", http.StatusBadRequest)
		return
	}

	dirName := r.FormValue("name")

	if dirName == "" {
		http.Error(w, "no directory name provided", http.StatusBadRequest)
		return
	} else if !validFilenameRegex.MatchString(dirName) {
		http.Error(w, "directory name contains illigal chars", http.StatusBadRequest)
		return
	}

	err := os.Mkdir(filepath.Join(parent, dirName), permDir)
	if err != nil && !os.IsExist(err) {
		http.Error(w, fmt.Sprintf("cludld not create %q", dirName), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, filepath.Join(cwd, dirName), http.StatusMovedPermanently)
}
