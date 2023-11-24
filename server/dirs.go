package reader

import (
	"os"
	"path/filepath"
)

func readDir(name string) (*Dir, error) {
	items, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	var dir Dir
	for _, item := range items {
		stat, err := os.Stat(filepath.Join(name, item.Name()))
		if err != nil {
			return nil, err
		}

		dir.Items = append(dir.Items, Item{
			Name:         item.Name(),
			LastModified: stat.ModTime(),
			IsDir:        stat.IsDir(),
		})
	}

	return &dir, nil
}

func isDir(name string) (bool, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}
