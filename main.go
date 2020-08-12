package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		name = flag.String("name", "sample.txt", "ファイル名")
	)
	flag.Parse()

	zipFile, err := os.Create(*name)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	a, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(a, strings.NewReader("1つめのファイルのテキストです。"))

	b, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(b, strings.NewReader("2つめのファイルのテキストです。"))

	fmt.Printf("created %s file", *name)
}
