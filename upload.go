// UPLOAD AND MKDIR
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	permFile = 0664
	permDir  = 0755
	// TODO:
	// add '.part' to extention to uncompleased files
	// upDefaultExt = ".part"
)

func (s *server) upload(w http.ResponseWriter, r *http.Request) {
	if dontAllowUploads {
		w.WriteHeader(http.StatusMethodNotAllowed)
		s.tmpl.ExecuteTemplate(w, "noup.html", nil)
		return
	}
	if r.Method != http.MethodPost && r.Method != http.MethodPatch && r.Method != http.MethodPut {
		_ = s.tmpl.ExecuteTemplate(w, "upload.html", nil)
		return
	}

	cwd, err := url.QueryUnescape(r.FormValue("cwd"))
	if cwd == "" || err != nil {
		http.Error(w, "no or bad cwd provided", http.StatusBadRequest)
		return
	}
	cwd = filepath.Join(s.root, strings.TrimPrefix(cwd, "/browse"))
	if stat, err := os.Stat(cwd); err != nil || !stat.IsDir() {
		http.Error(w, "bad file path", http.StatusBadRequest)
		return
	}

	fileName, err := url.QueryUnescape(r.FormValue("name"))
	if fileName == "" || err != nil {
		http.Error(w, "no name provided", http.StatusBadRequest)
		return
	}

	uuid := r.FormValue("uuid")
	if uuid == "" {
		http.Error(w, "no uuid provided", http.StatusBadRequest)
		return
	}
	fileWithUUID := filepath.Join(cwd, fileName+"_"+uuid)

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		http.Error(w, "offset err", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		file, err := os.Create(fileWithUUID)
		if err != nil {
			http.Error(w, "could not create file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
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

	} else { // put method
		sum := r.FormValue("sha256")
		if sum != "" {
			file, err := os.Open(fileWithUUID)
			if err != nil {
				http.Error(w, "something went wrong [0]", http.StatusBadRequest)
				return
			}

			hasher := sha256.New()
			if _, err := io.Copy(hasher, file); err != nil {
				file.Close()
				http.Error(w, "something went wrong [2]", http.StatusBadRequest)
				return
			}
			file.Close()

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
			return
		}
	}
}

var validFilenameRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9_\- .()+]+$`)

func (s *server) mkdir(w http.ResponseWriter, r *http.Request) {
	if dontAllowUploads {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	cwd := r.FormValue("cwd")
	if cwd == "" {
		http.Error(w, "cwd not provided", http.StatusBadRequest)
		return
	}

	parent := filepath.Join(s.root, strings.TrimPrefix(cwd, "/browse"))

	if pStat, err := os.Stat(parent); err != nil {
		http.Error(w, "could not resolve parent direcoty", http.StatusBadRequest)
		return
	} else if !pStat.IsDir() {
		http.Error(w, "paren is not a drectory", http.StatusBadRequest)
		return
	}

	dirName := r.FormValue("name")

	if dirName == "" {
		http.Error(w, "no directory name provided", http.StatusBadRequest)
		return
	} else if !validFilenameRegex.MatchString(dirName) {
		http.Error(w, "directory name contains illigal chars", http.StatusBadRequest)
		return
	}

	err := os.Mkdir(filepath.Join(parent, dirName), permDir)
	if err != nil && !os.IsExist(err) {
		http.Error(w, fmt.Sprintf("cludld not create %q", dirName), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("file creaded successfully"))
}
