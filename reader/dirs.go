package reader

import "os"

func ReadDir(name string) (*Dir, error) {
	items, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	var dir Dir
	for _, item := range items {
		stat, err := os.Stat(item.Name())
		if err != nil {
			return nil, err
		}

		dir = append(dir, Item{
			Name:         item.Name(),
			LastModified: stat.ModTime(),
			IsDir:        stat.IsDir(),
		})
	}

	return &dir, nil
}

func IsDir(name string) (bool, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}
