package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("VPS-helper")

	initHelper()

	//http.HandleFunc("/", hello)

	startBrowser("http://127.0.0.1:8080")

	http.ListenAndServe(":8080", nil)

}

func initHelper() {
	/*
		var fs http.FileSystem = assets

		file, err := fs.Open("./test.txt")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		data := make([]byte, 100)
		count, err := file.Read(data)
		fmt.Printf("Read %v Byte: %v\n", count, string(data[:count]))
	*/

	fs := http.FileServer(assets)
	http.Handle("/", http.StripPrefix("/", fs))
}

func generateStaticFiles() {
	cmd := exec.Command("./assets/assets.exe", "")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Generate Static Files Out: %s\n", out)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func startBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
