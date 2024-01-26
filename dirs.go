package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Name         string
	LastModified time.Time
	Size         string
	IsDir        bool
}

func (i Item) String() string {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "%s ", i.Name)

	return buff.String()
}

// readDir
//
// return the [dir, dir, dir, file, file, file]
func readDir(name string) ([]Item, error) {
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	var items []Item
	for _, entry := range entries {
		stat, err := entry.Info()
		if err != nil {
			return nil, err
		}

		items = append(items, Item{
			Name:         entry.Name(),
			LastModified: stat.ModTime(),
			Size:         prettySize(stat),
			IsDir:        stat.IsDir(),
		})
	}

	res := make([]Item, 0, len(items))

	for _, val := range items {
		if val.IsDir {
			res = append(res, val)
		}
	}

	for _, val := range items {
		if !val.IsDir {
			res = append(res, val)
		}
	}

	return res, nil
}

func prettySize(f os.FileInfo) string {
	if f.IsDir() {
		return "--"
	}
	s := float64(f.Size())

	switch {
	case s < 1024: // bytes
		return fmt.Sprintf("%.0f B", s)
	case s < 1024*1024: // kb
		return fmt.Sprintf("%.02f Kb", s/1024)
	case s < 1024*1024*1024: // mB
		return fmt.Sprintf("%.02f Mb", s/(1024*1024))
	default:
		return fmt.Sprintf("%.02f Gb", s/(1024*1024*1024))
	}
}
