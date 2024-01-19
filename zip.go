package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var zipFileNameCahce map[string]string = make(map[string]string)

func zipStr(val []string) string {
	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	reqFileNames := new(bytes.Buffer)
	for _, v := range val {
		reqFileNames.WriteByte(';')
		reqFileNames.WriteString(v)
	}

	sum := sha256.Sum256(reqFileNames.Bytes())
	return hex.EncodeToString(sum[:])
}

// download the zip of the given file
func (s *server) downZip(w http.ResponseWriter, r *http.Request) {
	if fileHash := strings.TrimPrefix(r.URL.Path, "/downzip/"); fileHash != "" {
		path, ok := zipFileNameCahce[fileHash]
		if !ok {
			http.Error(w, "could not find zip file", http.StatusBadRequest)
			return
		}
		http.ServeFile(w, r, path)
	}
}

// zip the files
func (s *server) zip(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	cwd := strings.TrimPrefix(r.FormValue("cwd"), "/browse")
	if cwd == "" {
		http.Error(w, "cwd not proved", http.StatusBadRequest)
		return
	}
	cwd = filepath.Join(s.root, cwd, "/")

	val, ok := r.Form["files"]
	if !ok {
		http.Error(w, "no files for zipping proved", http.StatusBadRequest)
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

	reqFileNamesHash := ""
	if len(val) == 1 {
		names := strings.Split(val[0], "/")
		nm := ""
		for i := len(names) - 1; i >= 0; i-- {
			if names[i] != "" {
				nm = names[i]
				break
			}
		}
		if nm == "browse" {
			nm = "root"
		}
		reqFileNamesHash = nm + ".zip"
	} else {
		reqFileNamesHash = zipStr(val) + ".zip"
	}

	if path, ok := zipFileNameCahce[reqFileNamesHash]; ok {
		fmt.Fprintf(w, "event: done\n")
		fmt.Fprintf(w, "data: "+`{"name": %q, "url": %q}`+"\n\n", reqFileNamesHash, path)
		flusher.Flush()
		return
	}

	res := []string{}
	for _, v := range val {
		if v = strings.TrimPrefix(v, "/browse/"); v == "" {
			continue
		}
		// NOTE: I don't know if it's a possibility for abbitray data access. so just incase.
		if strings.HasSuffix(v, "/..") && strings.HasPrefix(v, "../") && !strings.Contains(v, "/../") {
			http.Error(w, "Bad actor '..'", http.StatusBadRequest)
			return
		}
		res = append(res, filepath.Join(s.root, v))
	}

	callback := func(p int) error {
		_, err := fmt.Fprintf(w, "event: onProgress\ndata: {\"status\": %d}\n\n", p)
		if err != nil {
			return err
		}
		flusher.Flush()
		return nil
	}

	path, err := zipDirs(callback, s.zipSavePath, cwd, res...)
	if err != nil {
		log.Println("err while zipping:", err)
		fmt.Fprintf(w, "event: errror\ndata: {}\n\n")
		flusher.Flush()
		return
	}

	zipFileNameCahce[reqFileNamesHash] = path

	fmt.Fprintf(w, "event: done\n")
	fmt.Fprintf(w, "data: "+`{"name": %q, "url": %q}`+"\n\n",
		reqFileNamesHash, "/downzip/"+url.PathEscape(reqFileNamesHash))
	flusher.Flush()
}

func zipDirs(callback func(progress int) error, sDir, prefix string, dirs ...string) (string, error) {
	if len(prefix) > 0 && prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}
	var files []string
	for _, dir := range dirs {
		f, err := walkDirTree(dir)
		if err != nil {
			return "", err
		}
		files = append(files, f...)
	}

	r, err := os.Create(filepath.Join(sDir, fmt.Sprintf("%v.zip", time.Now().UnixMilli())))
	if err != nil {
		return "", err
	}
	defer r.Close()

	arc := zip.NewWriter(r)
	defer arc.Close()

	for i, f := range files {
		if err := callback(i * 100 / len(files)); err != nil {
			return "", err
		}

		r, err := os.Open(f)
		if err != nil {
			return "", err
		}

		w, err := arc.Create(strings.TrimPrefix(f, prefix))
		if err != nil {
			return "", err
		}
		_, err = io.Copy(w, r)
		if err != nil {
			return "", err
		}
		r.Close()
	}

	return r.Name(), nil
}

func walkDirTree(n string) ([]string, error) {
	if stat, err := os.Stat(n); err != nil {
		return nil, err
	} else if !stat.IsDir() {
		return []string{n}, nil
	}

	dirs, err := os.ReadDir(n)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, d := range dirs {
		if !d.IsDir() {
			paths = append(paths, filepath.Join(n, d.Name()))
			continue
		}

		p, err := walkDirTree(filepath.Join(n, d.Name()))
		if err != nil {
			return nil, err
		}
		paths = append(paths, p...)
	}

	return paths, nil
}
