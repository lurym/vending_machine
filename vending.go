package main

import (
    "net/http"
    "fmt"
	"encoding/json"
	"regexp"
	"errors"
)

func returnError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func snacks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/snacks called")
	if r.Method != http.MethodGet {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("snacks GET")
	json.NewEncoder(w).Encode(GetSnacksTable())
}


func coinSlots(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/coin_slots called")
	if r.Method != http.MethodPost {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("coin slots POST")
	slot := CreateNewCoinSlot()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(slot)
}

func UpdateCoinSlot(r *http.Request, slot CoinSlot) error {
	var c Coin
	err := json.NewDecoder(r.Body).Decode(&c)
	if err == nil {
		if c.Coin == 1 || c.Coin == 2 || c.Coin == 5 {
			fmt.Println("Updating balance")
			UpdateCoinSlotBalance(slot, c.Coin * 100)
			return nil
		} else {
			return errors.New("Invalid coin value")
		}
	}
	return err
}

func ValidateRequest(r *http.Request) (CoinSlot, error) {
	var slot CoinSlot
	re := regexp.MustCompile("/coin_slots/(\\w*)")
	submatch := re.FindStringSubmatch(r.URL.Path)
	if submatch != nil {
		slotID := submatch[1]
		fmt.Println("slotID: ", slotID)
		slot, keyExists := GetCoinSlot(slotID)
		if keyExists {
			return slot, nil
		}
		return slot, errors.New("Invalid slot ID")
	}
	return slot, errors.New("Invalid URL")
}

func singleCoinSlot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/coin_slots/:slot called")
	slot, err := ValidateRequest(r)
	if err != nil {
		fmt.Println("Error occured: ", err)
		returnError(w, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(slot)
	} else if r.Method == http.MethodPut {
		err := UpdateCoinSlot(r, slot)
		if err != nil {
			fmt.Println("Error occured: ", err)
			returnError(w, http.StatusBadRequest)
		}
	} else {
		returnError(w, http.StatusMethodNotAllowed)
	}
}

func main() {
	//FIXME it is better to use some ready URL router package, but I want to keep this sample as minimal as possible
    http.HandleFunc("/snacks", snacks)
    http.HandleFunc("/coin_slots", coinSlots)
    http.HandleFunc("/coin_slots/", singleCoinSlot)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Fatal error on ListenAndServe: ", err)
    }
}
