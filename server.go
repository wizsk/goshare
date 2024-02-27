package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	//go:embed frontend/src/*
	templateFiles embed.FS

	//go:embed static/*
	staticFiles embed.FS
)

type server struct {
	root     string
	tmp      string // tmp is the path where the zip files are stored
	showStat bool
	tmpl     tmplWrapper
	// zipped    map[string]string
}

// just a warapper for debugigg
type tmplWrapper interface {
	ExecuteTemplate(io.Writer, string, any) error
}

type tmpl struct{}

func (tp *tmpl) ExecuteTemplate(w io.Writer, name string, data any) error {
	t, err := newTmplate().ParseGlob("frontend/src/*")
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, name, data)
}

func newTmplate() *template.Template {
	return template.New("_base").Funcs(template.FuncMap{
		"pathJoin": func(base string, elem ...string) string {
			val, _ := url.JoinPath(base, elem...)
			return val
		},
		"timeFmt": func(t time.Time) string {
			return t.Format("01/02/2006 03:04 PM")
		},
	})
}

func newServer() server {
	if stat, err := os.Stat(rootDir); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("err: %q no such directory\n", rootDir)
		} else {
			fmt.Printf("err newServer: while opening the file: %v\n", err)
		}
		os.Exit(1)
	} else if !stat.IsDir() {
		fmt.Printf("err newServer: %q is not a directory", rootDir)
		os.Exit(1)
	}

	tmpDirPath, err := os.MkdirTemp(os.TempDir(), "goshare_zip_")
	if err != nil {
		log.Fatal(err)
	}

	var tr tmplWrapper
	if debug {
		tr = &tmpl{}
	} else {
		tr, err = newTmplate().ParseFS(templateFiles, "frontend/src/*")
		if err != nil {
			log.Fatal(err)
		}
	}

	return server{
		tmp:      tmpDirPath,
		root:     rootDir,
		showStat: !dontShowStat,
		tmpl:     tr,
	}
}

func (s *server) cleanup() {
	sighalChannel := make(chan os.Signal, 1)
	signal.Notify(sighalChannel, os.Interrupt, syscall.SIGTERM)
	<-sighalChannel

	fmt.Println("Cleaning up cache files")
	os.RemoveAll(s.tmp)
	os.Exit(0)
}

type svData struct {
	Dir  []Item // Directory
	Cd   string // current direcoty
	Od   string // working directory
	Umap []Umap
}

type Umap struct {
	Name, Url string
}

func (s *server) browse(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Join(s.root, strings.TrimPrefix(r.URL.Path, "/browse"))

	currentDirName := ""
	if stat, err := os.Stat(fileName); err != nil {
		s.tmpl.ExecuteTemplate(w, "filenotfound.html", nil)
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

	svd := svData{Od: r.URL.Path, Cd: currentDirName}

	var err error
	svd.Dir, err = readDir(fileName)
	if err != nil {
		log.Println(err)
		return
	}

	for _, itm := range strings.Split(r.URL.EscapedPath(), "/") {
		if len(itm) == 0 {
			continue
		}

		if i, err := url.QueryUnescape(itm); err == nil {
			itm = i
		}

		if len(svd.Umap) == 0 {
			svd.Umap = append(svd.Umap, Umap{itm, "/" + itm})
		} else {
			svd.Umap = append(svd.Umap, Umap{itm, svd.Umap[len(svd.Umap)-1].Url + "/" + itm})
		}

	}

	_ = s.tmpl.ExecuteTemplate(w, "index.html", &svd)
}

func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	path = filepath.Clean(path)
	http.ServeFileFS(w, r, staticFiles, path)
}
