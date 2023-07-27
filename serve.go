package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/wizsk/goshare/compress"
)

//go:embed tailwind/src/favicon.ico tailwind/src/output.css
var staticFiles embed.FS

func serveForm(w http.ResponseWriter, r *http.Request) {
	var data FormPageDatas
	if exisingRedirectURL := r.FormValue("redirect"); exisingRedirectURL != "" {
		data.RedirectURL = exisingRedirectURL
	} else {
		rawQ := ""
		if r.URL.RawQuery != "" {
			rawQ += "?" + r.URL.RawQuery

		}
		data.RedirectURL = url.QueryEscape(r.URL.Path + rawQ)
	}

	if indexTemplate.ExecuteTemplate(w, "form", data) != nil {
		http.Error(w, "someting went wrong", http.StatusInternalServerError)
	}
}

func serveResource(w http.ResponseWriter, file string) {
	switch file {
	case "css":
		css, err := staticFiles.ReadFile("tailwind/src/output.css")
		if err != nil {
			http.Error(w, "Failed to read css file", http.StatusInternalServerError)
			return
		}
		if styling {
			css, err = os.ReadFile("tailwind/src/output.css")
			if err != nil {
				http.Error(w, "Failed to read css file", http.StatusInternalServerError)
				return
			}
		}
		w.Header().Set("Content-Type", "text/css")
		w.Write(css)

	case "favicon":
		faviconData, err := staticFiles.ReadFile("tailwind/src/favicon.ico")
		if err != nil {
			http.Error(w, "Failed to read favicon.ico", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(faviconData)
	}
}

const ZIP_PATH = "/tmp/zipTest/"

func serveZipFile(w http.ResponseWriter, r *http.Request, zipType string) {
	if zipType == "down" {
		path := strings.TrimPrefix(r.URL.Path, "/")
		http.ServeFile(w, r, ZIP_PATH+filepath.Clean(path))
		return
	}

	// Set the response header for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// Flush the response to ensure the message is sent immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}
	flusher.Flush()

	fmt.Fprintf(w, "event: onProgress\n")
	d, _ := json.Marshal(ZipData{
		Status: "praparing",
	})
	fmt.Fprintf(w, "data: %s\n\n", d)
	flusher.Flush()

	fileUri := root + filepath.Clean(r.URL.Path)
	compress.ZIP_PATH = ZIP_PATH
	process := make(chan string)

	var file string
	var err error
	// ctx, cancel := context.WithCancel(context.Background())

	go func() {
		file, err = compress.Zip(r.Context(), fileUri, process)
	}()

	if err != nil {
		fmt.Fprintf(w, "event: errror\n")
		fmt.Fprintf(w, "data: {}\n\n")
		flusher.Flush()
		return
	}

	for val := range process {
		fmt.Fprintf(w, "event: onProgress\n")
		d, _ := json.Marshal(ZipData{
			Status: val,
		})
		fmt.Fprintf(w, "data: %s\n\n", d)
		flusher.Flush()
	}

	if err != nil {
		fmt.Fprintf(w, "event: errror\n")
		fmt.Fprintf(w, "data: {}\n\n")
	} else {
		fmt.Fprintf(w, "event: done\n")
		d, _ := json.Marshal(ZipData{
			Status: file,
		})
		fmt.Fprintf(w, "data: %s\n\n", d)
	}

	flusher.Flush()
}

type ZipData struct {
	Status string `json:"status"`
}
