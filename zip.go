package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var zipFileNameCahce map[string]string = make(map[string]string)

func walkDirTree(n string) ([]string, error) {
	n, err := filepath.EvalSymlinks(n)
	if err != nil {
		log.Fatal(err)
	}
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

func (s *server) zip(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	val, ok := r.Form["files"]
	if !ok {
		http.Error(w, "no files proved", http.StatusBadRequest)
		return
	}

	sort.Slice(val, func(i, j int) bool {
		return val[i] < val[j]
	})

	var reqFileNames strings.Builder
	for _, v := range val {
		reqFileNames.WriteByte(';')
		reqFileNames.WriteString(v)
	}

	if name, ok := zipFileNameCahce[reqFileNames.String()]; ok {
		fmt.Fprintf(w, "%q is file cache", name)
		return
	}

	res := []string{}
	for _, v := range val {
		if v = strings.TrimPrefix(strings.TrimSpace(v), "/browse/"); v == "" {
			continue
		}
		// NOTE: i don't know if it's a possibility
		if strings.HasSuffix(v, "/..") && strings.HasPrefix(v, "../") && !strings.Contains(v, "/../") {
			http.Error(w, "Bad actor '..'", http.StatusBadRequest)
			return
		}
		res = append(res, filepath.Join(s.root, v))
	}

	progress := make(chan int)
	var path string
	ctx, _ := context.WithCancel(context.Background())

	go func() {
		path, err = zipDirs(ctx, s.zipSavePath, progress, res...)
	}()

	for e := range progress {
		fmt.Println(e)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	zipFileNameCahce[reqFileNames.String()] = path
	fmt.Fprintf(w, "%q is writen", path)
}

func zipDirs(ctx context.Context, sDir string, progress chan<- int, dirs ...string) (string, error) {
	defer close(progress)
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

		w, err := arc.Create(f)
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
