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

	_, err = indexTemplate.New("cli").Parse(cliUiTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

func ServeWebUi(w http.ResponseWriter, r *http.Request) {
	var err error
	var datas Data
	datas.Directories, err = file(w, r)
	if err != nil {
		return
	}

	// web ui part
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

func file(w http.ResponseWriter, r *http.Request) ([]Directory, error) {
	var directories []Directory
	fileUri := root + filepath.Clean(r.URL.Path)
	file, err := os.Open(fileUri)
	if err != nil {
		http.Error(w, "File not found 404", http.StatusNotFound)
		log.Printf("root dir not found; %s\n", err)
		return directories, err
	}
	defer file.Close()

	dirs, err := file.ReadDir(0)
	if err != nil {
		http.ServeFile(w, r, fileUri)
		return directories, err
	}

	directories = directoriesList(dirs, fileUri, r.URL.Path)

	return directories, nil
}

func directoriesList(dirEntries []os.DirEntry, fileUri string, URLpath string) []Directory {
	var dirs []Directory
	for _, dir := range dirEntries {
		info, err := dir.Info()
		if err != nil {
			log.Println(err)
			continue
		}

		// -> / + "name" || /file + "/" + "name"
		path := URLpath
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

		dirs = append(dirs, dr)
	}

	return dirs
}
