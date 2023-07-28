package compress

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"
)

func getFilePaths(ctx context.Context, dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("canceled")
		default:
		}

		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func genRandomPostFix() string {
	return fmt.Sprint(time.Now().UnixNano())
}
