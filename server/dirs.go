package server

import (
	"fmt"
	"os"
)

func readDir(name string) (Dir, error) {
	items, err := os.ReadDir(name)
	if err != nil {
		return Dir{}, err
	}

	var dir Dir
	for _, item := range items {
		stat, err := item.Info()
		if err != nil {
			return Dir{}, err
		}

		dir.Items = append(dir.Items, Item{
			Name:         item.Name(),
			LastModified: stat.ModTime(),
			Size:         prettySize(stat),
			IsDir:        stat.IsDir(),
		})
	}

	return dir, nil
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
