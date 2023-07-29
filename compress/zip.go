package compress

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

var ZIP_PATH string

type ZipFileInfo struct {
	Name     string
	FilePath string
}

func Zip(ctx context.Context, progress chan<- string, rootDir string) (ZipFileInfo, error) {
	defer close(progress)

	var output ZipFileInfo
	var outputFileName string
	var trimFromRoot string

	if stat, err := os.Stat(rootDir); err != nil {
		return output, err
	} else {
		outputFileName = stat.Name()
		trimFromRoot = strings.TrimRight(rootDir, outputFileName)
	}

	if ZIP_PATH == "" {
		return output, fmt.Errorf("please specify the ZIP_PATH")
	}
	// else if ZIP_PATH[len(ZIP_PATH)-1] != '/' {
	// 	ZIP_PATH += "/"
	// }

	progress <- "preparing"
	files, err := getFilePaths(ctx, rootDir)
	if err != nil {
		return output, err
	}

	filesCount := float64(len(files))

	// je name dia create kori file.Name() dakle oi name e dei!!
	archive, err := os.Create(fmt.Sprintf("%s/%s.zip-%s", ZIP_PATH, outputFileName, genRandomPostFix()))
	if err != nil {
		return output, err
	}
	defer archive.Close()

	zipper := zip.NewWriter(archive)
	defer zipper.Close()

zipLoop:
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
			err = fmt.Errorf("cancelled by context")
			break zipLoop
		default:
		}
	}
	if err != nil {
		return output, err
	}

	output.FilePath = archive.Name()
	output.Name = outputFileName + ".zip"
	return output, nil
}
