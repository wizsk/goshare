package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	permFile = 0664
	permDir  = 0755

	upDefaultExt = ".part"
)

func (s *server) upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPatch && r.Method != http.MethodPut {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	cwd := r.FormValue("cwd")
	if cwd == "" {
		http.Error(w, "no cwd provided", http.StatusBadRequest)
		return
	}
	cwd = filepath.Join(s.root, strings.TrimPrefix(cwd, "/browse"))
	if stat, err := os.Stat(cwd); err != nil || !stat.IsDir() {
		http.Error(w, "bad file path", http.StatusBadRequest)
		return
	}

	file := r.FormValue("name")
	if file == "" {
		http.Error(w, "no name provided", http.StatusBadRequest)
		return
	}
	file = filepath.Join(cwd, file)

	uuid := r.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "no uuid provided", http.StatusBadRequest)
		return
	}

	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		http.Error(w, "size err", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		http.Error(w, "offset err", http.StatusBadRequest)
		return
	}

	fmt.Println(r.Method, cwd, file, uuid, size, offset)

	// no proxy or anything will be used so headers will be used to communicate aobut files
	if r.Method == http.MethodPost {
		_, err := os.Create(file)
		if err != nil {
			http.Error(w, "could not create file", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == http.MethodPut {
		file, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, permFile)
		if err != nil {
			http.Error(w, "could not open file", http.StatusBadRequest)
			return
		}
		fileStat, err := file.Stat()
		if err != nil {
			http.Error(w, "could not get stat of file", http.StatusInternalServerError)
			return
		}
		if offset != int(fileStat.Size()) {
			http.Error(w, "file corrupted", http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(w, r.Body); err != nil {
			http.Error(w, "couldn't write file", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		// this was already checked
	}
}
