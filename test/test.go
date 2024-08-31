package main

import (
	"fmt"
	"os"
	core "peer-store/core/file_manager"
	"peer-store/core/pki"
)

func main() {
	pki_test()
}

func pki_test() {
	pki.Generate_pki_key_pair()
}

func file_chunk_test() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic("Error while reading the file path. Please check your file path")
	}
	baseDir := currentDir + "/test/"
	file_path_string := baseDir + "test.pdf"
	l, err := core.FileSpliterService(file_path_string, 10, baseDir)

	if err != nil {
		fmt.Println("Error while chunking files : " + err.Error())
	}

	fmt.Println(l)
}
