package zipFunctions

import (
	"archive/zip"
	"bytes"
	"fmt"
	generics "golang_test/genericFunctions"
	"golang_test/model"
	"io/ioutil"
)

// ZipData This function create a buffered zip file with MyFiles content
func ZipData(files *[]model.MyFiles) (bufferedZip []byte) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)

	// Fill the zip archive with files variable content
	for _, file := range *files {
		zipFile, err := zipWriter.Create(file.Name)
		generics.Check(err)
		_, err = zipFile.Write([]byte(file.Body))
		generics.Check(err)
	}
	// Close the zip archive
	generics.CloseFile(zipWriter)

	// return the buffered zip file
	return buf.Bytes()

}

// ReadAll function read all content in the zipped File
func readAll(file *zip.File) []byte {
	fc, err := file.Open()
	generics.Check(err)
	defer generics.CloseFile(fc)

	content, err := ioutil.ReadAll(fc)
	generics.Check(err)

	return content
}

// ReadAllElements function read all files in the archive
func ReadAllElements(archive **zip.Reader) {
	for _, file := range (*archive).File {
		myFile := readAll(file)
		fmt.Printf("File %s has Body:\n %s\n\n", file.Name, myFile)
	}
}
