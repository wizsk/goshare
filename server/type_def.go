package server

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"time"
)

type Server struct {
	root, tmp string
	tmpl      *template.Template
}

func NewServer(r, t, templateDir string) *Server {
	sv := Server{root: r, tmp: t}
	var err error
	sv.tmpl, err = template.ParseGlob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		log.Fatal(err)
	}

}

type Dir struct {
	Items []Item
}

type Item struct {
	Name         string
	LastModified time.Time
	Size         string
	IsDir        bool
}

func (d *Dir) String() string {
	buff := new(bytes.Buffer)

	for i, itm := range d.Items {
		// if i != 0 {
		// 	fmt.Fprintf(buff, "\nname: %s, isDir: %v, lastModified: %v", itm.Name, itm.IsDir, itm.LastModified)
		// } else {
		// 	fmt.Fprintf(buff, "name: %s, isDir: %v, lastModified: %v", itm.Name, itm.IsDir, itm.LastModified)
		// }
		if i != 0 {
			fmt.Fprintf(buff, "\n%s", itm.Name)
		} else {
			fmt.Fprintf(buff, "%s", itm.Name)
		}
	}

	return buff.String()
}
