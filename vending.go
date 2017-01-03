package main

import (
    "net/http"
    "fmt"
	"encoding/json"
)

func returnError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func snacksGET(w http.ResponseWriter, r *http.Request) {
	fmt.Println("snacks GET")
	json.NewEncoder(w).Encode(snacksTable)
}

func snacks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/snacks called")
	if r.Method != http.MethodGet {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	snacksGET(w, r)
}


func coinSlotsPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("coin slots POST")
	
}

func coinSlots(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/coin_slots called")
	if r.Method != http.MethodPost {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	coinSlotsPOST(w, r)
}

func main() {
    http.HandleFunc("/snacks", snacks)
    http.HandleFunc("/coin_slots", coinSlots)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Fatal error on ListenAndServe: ", err)
    }
}
