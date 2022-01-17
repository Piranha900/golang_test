package main

// https://gist.github.com/madevelopers/40b269730df687cdcb8b

import (
	"archive/zip"
	"bytes"
	"golang_test/genericFunctions"
	"golang_test/model"
	"golang_test/zipFunctions"
)

var files = []model.MyFiles{
	{"readme.txt", "This archive contains some text files."},
	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	{"todo.txt", "Get animal handling licence.\nWrite more examples."},
}

func main() {
	byteArchive := zipFunctions.ZipData(&files)

	zipArchive, err := zip.NewReader(bytes.NewReader(byteArchive), bytes.NewReader(byteArchive).Size())
	genericFunctions.Check(err)

	zipFunctions.ReadAllElements(&zipArchive)

}
