package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

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

func (s *server) zipDirs(dirs ...string) (string, error) {
	var files []string
	for _, dir := range dirs {
		f, err := walkDirTree(dir)
		if err != nil {
			return "", err
		}
		files = append(files, f...)
	}

	r, err := os.Create(filepath.Join(s.tmp, fmt.Sprintf("%v.zip", time.Now().UnixMilli())))
	if err != nil {
		return "", err
	}
	defer r.Close()

	arc := zip.NewWriter(r)
	defer arc.Close()
	for _, f := range files {
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
	}

	return r.Name(), nil
}
