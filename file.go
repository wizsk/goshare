package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

//go:embed tailwind/src/index.html tailwind/src/form.html tailwind/src/index/*
var templateFiles embed.FS
var indexTemplate *template.Template

type Directory struct {
	Name      string
	SizeBytes int64
	Size      string
	Url       string
	IsDir     bool
	Icon      template.HTML
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

	indexTemplate, err = template.ParseFS(templateFiles,
		"tailwind/src/index.html",
		"tailwind/src/form.html",
		"tailwind/src/index/components.html",
		"tailwind/src/index/list.html",
		"tailwind/src/index/index.js",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = indexTemplate.New("cli").Parse(cliUiTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

func file(w http.ResponseWriter, r *http.Request) ([]Directory, error) {
	var directories []Directory
	fileUri := root + filepath.Clean(r.URL.Path)
	file, err := os.Open(fileUri)
	if err != nil {
		http.Error(w, "File not found 404", http.StatusNotFound)
		// log.Printf("root dir not found; %s\n", err)
		return directories, err
	}
	defer file.Close()

	dirs, err := file.ReadDir(0)
	if err != nil {
		printStat(r, FILE_DOWN)
		http.ServeFile(w, r, fileUri)
		return directories, err
	}

	directories = directoriesList(dirs, r)
	sortDir(directories, r.FormValue("sort"))

	return directories, nil
}

func directoriesList(dirEntries []os.DirEntry, r *http.Request) []Directory {
	// fileUri := root + filepath.Clean(r.URL.Path)
	var dirs []Directory
	for _, dir := range dirEntries {
		info, err := dir.Info()
		if err != nil {
			log.Println(err)
			continue
		}

		// -> / + "name" || /file + "/" + "name"
		path := r.URL.Path
		if path == "/" {
			path += url.PathEscape(dir.Name())
		} else {
			path += "/" + url.PathEscape(dir.Name())
		}

		dr := Directory{
			Name:      dir.Name(),
			SizeBytes: info.Size(),
			Size:      fileSize(info.Size()),
			Url:       path,
			IsDir:     dir.IsDir(),
			Icon:      directoryIcon,
		}

		// if !dir.IsDir() {
		// 	dr.Icon = detectFileType(fileUri + "/" + dir.Name())
		// }

		dirs = append(dirs, dr)
	}

	return dirs
}
