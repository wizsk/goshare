package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

func fileSize(size int64) string {
	s := float64(size)
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

func sortDir(d []Directory, item string) {
	switch item {
	case "namedesc":
		sort.Slice(d, func(i, j int) bool {
			return d[i].Name > d[j].Name
		})
	case "sizeasc":
		sort.Slice(d, func(i, j int) bool {
			return d[i].SizeBytes < d[j].SizeBytes
		})
	case "sizedesc":
		sort.Slice(d, func(i, j int) bool {
			return d[i].SizeBytes > d[j].SizeBytes
		})
	case "bydir":
		sort.Slice(d, func(i, j int) bool {
			if d[i].IsDir && d[j].IsDir {
				return false
			}
			if d[i].IsDir {
				return true
			}
			// default d[j].IsDir
			return false
		})
	case "byfile":
		sort.Slice(d, func(i, j int) bool {
			if d[i].IsDir && d[j].IsDir {
				return false
			}
			if !d[i].IsDir {
				return true
			}
			// defaulty !d[j].IsDir
			return false
		})
	default:
		// "nameasc"
		sort.Slice(d, func(i, j int) bool {
			return d[i].Name < d[j].Name
		})
	}
}

/*
const (
	// ANSI escape code for resetting text color to default
	CONSOL_Reset = "\033[0m"

	// ANSI escape codes for different colors with CONSOL_ prefix
	CONSOL_Red    = "\033[31m"
	CONSOL_Green  = "\033[32m"
	CONSOL_Yellow = "\033[33m"
	CONSOL_Blue   = "\033[34m"
	CONSOL_Purple = "\033[35m"
	CONSOL_Cyan   = "\033[36m"
	CONSOL_White  = "\033[37m"
)
*/

// print the pretty output ^_^
func printStat(r *http.Request, from string) {
	// formPadding := (18 - len(from)) / 2
	// addr := strings.Split(r.RemoteAddr, ":")[0]
	// addrPadding := (17 - len(addr)) / 2

	fmt.Printf("[%16s] %s | %15s | %s | %q\n",
		from,
		time.Now().Format("2006/01/02 - 03:04:05 PM"),
		strings.Split(r.RemoteAddr, ":")[0],
		r.Method,
		r.URL.Path,
	)
	// fmt.Printf("[%*s%s%*s] %s | %*s%s%*s | %s | %q\n",
	// 	// formPadding, "", from, formPadding, "",
	// 	time.Now().Format("2006/01/02 - 03:04:05 PM"),
	// 	// addrPadding, "", addr, addrPadding, "",
	// 	r.Method,
	// 	r.URL.Path,
	// )
}

// for genarating root/file/..
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
