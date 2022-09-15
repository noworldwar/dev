package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		outputHTML(w, r, "index.html")
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		outputHTML(w, r, "login.html")
	})

	fmt.Println(http.ListenAndServe(":19094", r))

	fmt.Println("Client is running at 19094 port.")
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
