package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//go:embed tailwind/src/index.html
var templateFiles embed.FS
var indexTemplate *template.Template

type ProgessPah struct {
	Title    string
	Url      string
	SlashPre bool
}

type Data struct {
	Dirtitle     string
	PreviousPage string
	Directories  []Directory
	ProgessPah   []ProgessPah
}

type Directory struct {
	Name  string
	Size  string
	Url   string
	IsDir bool
	Icon  template.HTML
}

var root string

func fileSeverInit(file string) {
	if file == "" {
		log.Fatal("root file should not be empty")
	}
	root = file
	rootFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer rootFile.Close()

	_, err = rootFile.ReadDir(1)
	if err != nil {
		log.Fatal("root file should be a directory;", err)
	}

	indexTemplate, err = template.ParseFS(templateFiles, "tailwind/src/index.html")
	if err != nil {
		log.Fatal(err)
	}

}

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	fileUri := root + filepath.Clean(r.URL.Path)
	file, err := os.Open(fileUri)
	if err != nil {
		http.Error(w, "File not found 404", http.StatusNotFound)
		log.Printf("root dir not found; %s\n", err)
		return
	}
	defer file.Close()

	dirs, err := file.ReadDir(0)
	if err != nil {
		http.ServeFile(w, r, fileUri)
		return
	}

	var datas Data

	for _, dir := range dirs {
		info, err := dir.Info()
		if err != nil {
			log.Println(err)
			continue
		}

		path := r.URL.Path
		if path == "/" {
			path += url.PathEscape(dir.Name())
		} else {
			path += "/" + url.PathEscape(dir.Name())
		}

		dr := Directory{
			Name:  dir.Name(),
			Size:  fileSize(info),
			Url:   path,
			IsDir: dir.IsDir(),
			Icon:  directoryIcon,
		}

		if !dir.IsDir() {
			dr.Icon = detectFileType(fileUri + "/" + dir.Name())
		}

		datas.Directories = append(datas.Directories, dr)
	}

	if split := strings.Split(r.URL.EscapedPath(), "/"); len(split) > 2 {
		datas.PreviousPage = strings.Join(split[:len(split)-1], "/")
	} else {
		datas.PreviousPage = "/"
	}

	datas.ProgessPah = possiblePahts(r)
	datas.Dirtitle = datas.ProgessPah[len(datas.ProgessPah)-1].Title

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = indexTemplate.ExecuteTemplate(w, "main", datas)
	if err != nil {
		log.Println(err)
		return
	}
}
