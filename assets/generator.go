// vfsgen.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var vfsdata http.FileSystem = http.Dir("vfsdata")
	err := vfsgen.Generate(vfsdata, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}

	count, err := copyFile("../assets_vfsdata.go", "./assets_vfsdata.go")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The assets_vfsdata.go is %d KB\n", count/1024)
}

func copyFile(dstFile, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
