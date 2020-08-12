package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		copySource      = flag.String("copySource", "copySource.txt", "コピー元ファイル名")
		copyDestination = flag.String("copyDestination", "copyDestination.txt", "コピー先ファイル名")
	)
	flag.Parse()

	oldFile, err := os.Open(*copySource)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	newFile, err := os.Create(*copyDestination)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)

	fmt.Printf("finished copy %s to %s", *copySource, *copyDestination)
}
