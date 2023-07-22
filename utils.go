package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/wizsk/goshare/auth"
)

const usages = `
USAGE:
    goshare [OPTIONS]

OPTIONS:
    -d
    	Direcotry path
    -p
        Password (default is none)
    -port 
        Port number (default is "8001")
`

func authHelper(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad req", http.StatusBadRequest)
		return
	}
	inputPass := r.FormValue("password")
	if *pass == inputPass {
		auth.WriteCookie(w)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func fileSize(file fs.FileInfo) string {
	s := float64(file.Size())
	switch {
	case s < 1024:
		return fmt.Sprintf("%.0f B", s)
	case s < 1024*1024:
		return fmt.Sprintf("%.01f Kb", s/1024)
	case s < 1024*1024*1024:
		return fmt.Sprintf("%.01f Mb", (s/1024)/1024)
	case s >= 1024*1024*1024:
		return fmt.Sprintf("%.01f Gb", ((s/1024)/1024)/1024)
	}

	return ""
}

func possiblePahts(r *http.Request) []ProgessPah {
	var p []ProgessPah
	poosiblePaht := ""
	for i, v := range strings.Split(strings.TrimRight(r.URL.EscapedPath(), "/"), "/") {
		if v == "" {
			p = append(p, ProgessPah{
				Title:    "root/",
				Url:      "/",
				SlashPre: false,
			})
			continue
		}

		poosiblePaht += "/" + v
		title, _ := url.PathUnescape(v)
		p = append(p, ProgessPah{
			Title:    title,
			Url:      poosiblePaht,
			SlashPre: true,
		})
		if i == 1 {
			p[i].SlashPre = false
		}
	}
	return p
}

func detectFileType(filePath string) template.HTML {
	file, err := os.Open(filePath)
	if err != nil {
		return unknownFileIcon
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return unknownFileIcon
	}

	fileType := http.DetectContentType(buffer)

	switch fileType {
	case "image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp",
		"image/tiff", "image/x-icon", "image/svg+xml", "image/vnd.adobe.photoshop":
		return imgIcon
	case "video/mp4", "video/quicktime", "video/x-msvideo", "video/x-matroska",
		"video/webm", "video/x-flv", "video/3gpp":
		return videoIcon
	case "audio/mpeg", "audio/wav", "audio/midi", "audio/ogg", "audio/x-flac",
		"audio/x-ms-wma", "audio/x-musepack", "audio/vnd.rn-realaudio", "audio/webm":
		return audioIcon
	case "application/pdf":
		return pdfIcon
	case "text/plain", "text/html", "text/xml", "application/json", "application/xml",
		"application/x-yaml", "text/csv":
		return textIcon
	case "application/zip", "application/x-tar", "application/x-gzip", "application/x-bzip2", "application/x-rar-compressed",
		"application/x-7z-compressed", "application/x-zip-compressed", "application/x-java-archive":
		return archiveIcon
	case "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-powerpoint", "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"application/vnd.oasis.opendocument.text", "application/vnd.oasis.opendocument.spreadsheet", "application/vnd.oasis.opendocument.presentation":
		return documentIcon
	default:
		return unknownFileIcon
	}
}

/*
func detectFileType(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "Unknown"
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "Unknown"
	}

	fileType := http.DetectContentType(buffer)

	switch fileType {
	case "image/jpeg", "image/png":
		return "Image"
	case "video/mp4", "video/quicktime", "video/x-msvideo", "video/x-matroska":
		return "Video"
	case "audio/mpeg", "audio/wav", "audio/midi":
		return "Audio"
	case "application/pdf":
		return "PDF Document"
	case "text/plain":
		return "Text File"
	case "application/zip", "application/x-tar", "application/x-gzip", "application/x-bzip2", "application/x-rar-compressed":
		return "Archive"
	case "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.ms-excel", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/vnd.ms-powerpoint", "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return "Document"
	default:
		return "Unknown"
	}
}
*/
