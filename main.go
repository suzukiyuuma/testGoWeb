package main

import (
	"fmt"
    "net/http"
    "encoding/json"
)

type User struct {
    Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    name := r.FormValue("name")
    if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
        return
    } 

	user := User{
		Name: name,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(user)

}

func main() {
	fmt.Println("Webサーバー立ち上げ開始")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}