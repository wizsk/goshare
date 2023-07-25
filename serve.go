package main

import (
	"embed"
	"net/http"
	"net/url"
)

//go:embed tailwind/src/favicon.ico tailwind/src/output.css
var staticFiles embed.FS

func serveForm(w http.ResponseWriter, r *http.Request) {
	data := FormPageDatas{
		RedirectURL: url.QueryEscape(r.URL.Path + "?" + r.URL.RawQuery),
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
