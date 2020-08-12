package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
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
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
