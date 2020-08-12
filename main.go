package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		name = flag.String("name", "sample.txt", "ファイル名")
	)
	flag.Parse()

	file, err := os.Create(*name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)

	fmt.Printf("created %s file", *name)
}
