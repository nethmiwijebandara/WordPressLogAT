package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(r *http.Request) ([]string, error) {

	file, handler, err := r.FormFile("zip") // zip is the key of the form-data
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(handler.Filename)
	fmt.Println()
	fmt.Println(handler.Header)
	fileSize, err := file.Seek(0, 2)

	var filenames []string
	read, err := zip.NewReader(file, fileSize)

	if err != nil {
		return filenames, err
	}

	// Executes until files in the source directory keep storing filenames and extracts into destination until an error
	for _, f := range read.File {

		// Store "path/filename" for returning and using later on
		fpath := filepath.Join("unzipped", f.Name)
		fmt.Println(fpath)
		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean("unzipped")+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		//append the accessed filenames
		filenames = append(filenames, fpath)
		fmt.Println(filenames)
		if f.FileInfo().IsDir() {

			// Creating a new Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())

		if err != nil {
			return filenames, err
		}
		rc, err := f.Open()

		if err != nil {
			return filenames, err
		}
		_, err = io.Copy(outFile, rc)

		//close the file without defer
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
