package compress

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

var ZIP_PATH = os.TempDir()

func Zip(ctx context.Context, rootDir string, progress chan<- string) (string, error) {
	defer close(progress)

	var outputFileName string
	var trimFromRoot string

	if stat, err := os.Stat(rootDir); err != nil {
		return "", err
	} else {
		outputFileName = stat.Name()
		trimFromRoot = strings.TrimRight(rootDir, outputFileName)
	}

	if ZIP_PATH == "" {
		return "", fmt.Errorf("please specify the ZIP_PATH")
	} else if ZIP_PATH[len(ZIP_PATH)-1] != '/' {
		ZIP_PATH += "/"
	}

	progress <- "preparing"
	files, err := getFilePaths(ctx, rootDir)
	if err != nil {
		return "", err
	}

	filesCount := float64(len(files))

	archive, err := os.Create(ZIP_PATH + outputFileName + ".zip")
	if err != nil {
		return "", err
	}
	defer archive.Close()

	zipper := zip.NewWriter(archive)
	defer zipper.Close()

	progress <- ""
	for i, file := range files {
		var opneFile *os.File
		opneFile, err = os.Open(file)
		if err != nil {
			break
		}

		var zw io.Writer
		zw, err = zipper.Create(strings.TrimPrefix(file, trimFromRoot))
		if err != nil {
			break
		}

		if _, err = io.Copy(zw, opneFile); err != nil {
			break
		}

		opneFile.Close()
		progress <- fmt.Sprintf("zipping %0.1f%%", (float64(i)/filesCount)*100)
		select {
		case <-ctx.Done():
			fmt.Println("cancel hoiye gese")
			return "", fmt.Errorf("cancelled")
		default:
			continue
		}
	}
	if err != nil {
		return "", err
	}

	return outputFileName + ".zip", nil
}
