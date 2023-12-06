package main

import (
	"crypto/sha256"
	"encoding/hex"
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
	fileWithUUID := filepath.Join(cwd, fileName+"_"+uuid)

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
		file, err := os.Create(fileWithUUID)
		if err != nil {
			http.Error(w, "could not create file", http.StatusInternalServerError)
			return
		}
		file.Close()
		w.WriteHeader(http.StatusCreated)

	} else if r.Method == http.MethodPatch {
		file, err := os.OpenFile(fileWithUUID, os.O_APPEND|os.O_WRONLY, permFile)
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

		if offset != int(fileStat.Size()) {
			http.Error(w, "file corrupted", http.StatusInternalServerError)
			return
		}

		if _, err = io.Copy(file, r.Body); err != nil {
			http.Error(w, "couldn't write file", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		// this was already checked
	} else {
		// put method
		sum := r.FormValue("sha256")
		if sum != "" {
			file, err := os.Open(fileWithUUID)
			if err != nil {
				http.Error(w, "something went wrong [0]", http.StatusBadRequest)
				return
			}
			defer file.Close()

			hasher := sha256.New()
			if _, err := io.Copy(hasher, file); err != nil {
				http.Error(w, "something went wrong [2]", http.StatusBadRequest)
				return
			}

			if gotSum := hex.EncodeToString(hasher.Sum(nil)); gotSum != sum {
				fmt.Println(fileName, sum, gotSum)
				http.Error(w, "sum don't match", http.StatusBadRequest)
				return
			}
		}

		ext := filepath.Ext(fileName)
		rawFileName := filepath.Join(cwd, strings.TrimSuffix(fileName, ext))
		add := ""

		for i := 1; i < 100; i++ {
			if _, err := os.Stat(rawFileName + add + ext); os.IsNotExist(err) {
				break
			}
			add = "-" + strconv.Itoa(i)
		}
		if err := os.Rename(fileWithUUID, rawFileName+add+ext); err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
