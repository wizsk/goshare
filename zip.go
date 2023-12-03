package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var zipFileNameCahce map[string]string = make(map[string]string)

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

	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	var reqFileNames strings.Builder
	for _, v := range val {
		reqFileNames.WriteByte(';')
		reqFileNames.WriteString(v)
	}

	if path, ok := zipFileNameCahce[reqFileNames.String()]; ok {
		fmt.Fprintf(w, "event: done\n")
		fmt.Fprintf(w, "data: "+`{"name": %q, "url": %q}`+"\n\n", path, reqFileNames.String())
		flusher.Flush()
		return
	}

	res := []string{}
	for _, v := range val {
		if v = strings.TrimPrefix(strings.TrimSpace(v), "/browse/"); v == "" {
			continue
		}
		// NOTE: i don't know if it's a possibility for abbitray data access. so just incase.
		if strings.HasSuffix(v, "/..") && strings.HasPrefix(v, "../") && !strings.Contains(v, "/../") {
			http.Error(w, "Bad actor '..'", http.StatusBadRequest)
			return
		}
		res = append(res, filepath.Join(s.root, v))
	}

	progress := make(chan int)
	var path string
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancle will be called automatically

	go func() {
		path, err = zipDirs(ctx, s.zipSavePath, cwd, progress, res...)
	}()

	for e := range progress {
		_, err := fmt.Fprintf(w, "event: onProgress\ndata: {\"status\": %d}\n\n", e)
		if err != nil {
			return
		}
		flusher.Flush()
	}

	if err != nil {
		fmt.Fprintf(w, "event: errror\n")
		fmt.Fprintf(w, "data: {}\n\n")
		fmt.Println(err)
		flusher.Flush()
		return
	}

	zipFileNameCahce[reqFileNames.String()] = path

	fmt.Fprintf(w, "event: done\n")
	fmt.Fprintf(w, "data: "+`{"name": %q, "url": %q}`+"\n\n", path, reqFileNames.String())
	flusher.Flush()
	// fmt.Fprintf(w, "%q is writen", path)
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

func zipDirs(ctx context.Context, sDir, prefix string, progress chan<- int, dirs ...string) (string, error) {
	defer close(progress)

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
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("canceled")
		default:
		}

		progress <- i * 100 / len(files)

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
