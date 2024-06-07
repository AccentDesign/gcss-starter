package main

import (
	"fmt"
	"gcss.starter.dev/styles"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	dir        = currentDirectory()
	stylesheet = styles.NewStyleSheet()
)

func currentDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return filepath.Dir(filename)
}

func main() {
	asFile()
	asServer()
}

func asFile() {
	file, err := os.Create("tmp/stylesheet.css")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := stylesheet.CSS(file); err != nil {
		panic(err)
	}
}

func asServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(dir, "index.html"))
	})
	http.HandleFunc("/stylesheet.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		if err := stylesheet.CSS(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
