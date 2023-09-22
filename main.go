package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatal(err)
	}
	if err := html.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
	// fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/view", viewHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("192.168.33.10:8080", nil))
}