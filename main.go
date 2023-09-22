package main

import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Hello World")
}
func main(){
	fmt.Println("webサーバー立ち上げ開始")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}