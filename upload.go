package main

import (
	"fmt"
	"io"
	"log"
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

	fileName := r.FormValue("name")
	if fileName == "" {
		http.Error(w, "no name provided", http.StatusBadRequest)
		return
	}

	uuid := r.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "no uuid provided", http.StatusBadRequest)
		return
	}
	fileUUID := filepath.Join(cwd, fileName+"_"+uuid)

	// size, err := strconv.Atoi(r.FormValue("size"))
	// if err != nil {
	// 	http.Error(w, "size err", http.StatusBadRequest)
	// 	return
	// }

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		http.Error(w, "offset err", http.StatusBadRequest)
		return
	}

	// fmt.Println(r.Method, cwd, file, uuid, size, offset)

	// no proxy or anything will be used so headers will be used to communicate aobut files
	if r.Method == http.MethodPost {
		file, err := os.Create(fileUUID)
		if err != nil {
			http.Error(w, "could not create file", http.StatusInternalServerError)
			return
		}
		file.Close()
		w.WriteHeader(http.StatusCreated)

	} else if r.Method == http.MethodPatch {
		file, err := os.OpenFile(fileUUID, os.O_APPEND|os.O_WRONLY, permFile)
		if err != nil {
			http.Error(w, "could not open file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		fileStat, err := file.Stat()
		if err != nil {
			http.Error(w, "could not get stat of file", http.StatusInternalServerError)
			return
		}

		fmt.Println("file size", offset, fileStat.Size())
		if offset != int(fileStat.Size()) {
			http.Error(w, "file corrupted", http.StatusInternalServerError)
			return
		}

		var written int64
		if written, err = io.Copy(file, r.Body); err != nil {
			http.Error(w, "couldn't write file", http.StatusInternalServerError)
			return
		}
		r.Body.Close()
		fmt.Println("wrritten", written)
		// this was already checked
	} else {
		sum := r.FormValue("sha256")
		_ = sum
		if err != nil {
			http.Error(w, "offset err", http.StatusBadRequest)
			return
		}
		// put
		rawFileName := filepath.Join(cwd, fileName)
		add := ""
		for i := 1; i < 100; i++ {
			if _, err := os.Stat(rawFileName + add); os.IsNotExist(err) {
				break
			}
			add = "." + strconv.Itoa(i)
		}
		if err := os.Rename(fileUUID, rawFileName+add); err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
