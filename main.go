package main

import "fmt"
import "net/http"

const baseMessage = "Hello World"

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    if name == "" {
        fmt.Fprint(w, baseMessage)
        return
    }

    fmt.Fprintf(w, "%s, %s", baseMessage, name)
}

func main(){
	fmt.Println("webサーバー立ち上げ開始")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}