package main

import (
    "net/http"
    "fmt"
)

func returnError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func snacks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/snacks called")
	if r.Method != http.MethodGet {
		returnError(w, http.StatusMethodNotAllowed)
	}
	fmt.Println("GET")
	
    fmt.Fprintf(w, "Hello world!") 
}

func main() {
    http.HandleFunc("/snacks", snacks)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Fatal error on ListenAndServe: ", err)
    }
}
