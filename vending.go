package main

import (
    "net/http"
    "fmt"
	"encoding/json"
	"regexp"
	"errors"
	"strconv"
)

func returnError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func Snacks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/snacks called")
	if r.Method != http.MethodGet {
		returnError(w, http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("snacks GET")
	json.NewEncoder(w).Encode(GetSnacksTable())
}


func CoinSlots(w http.ResponseWriter, r *http.Request) {
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
		}
		return errors.New("Invalid coin value")
	}
	return err
}

func ValidateRequestUpdateSlotCoin(r *http.Request) (CoinSlot, error) {
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

func SingleCoinSlot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/coin_slots/:slot called")
	slot, err := ValidateRequestUpdateSlotCoin(r)
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

func ValidateRequestBuySnack(r *http.Request) (CoinSlot, Snack, error) {
	var slot CoinSlot
	var snack Snack
	re := regexp.MustCompile("/snacks/(\\d*)")
	submatch := re.FindStringSubmatch(r.URL.Path)
	if submatch == nil {
		return slot, snack, errors.New("Invalid URL")
	}

	snackID := submatch[1]
	fmt.Println("snackID: ", snackID)
	snackIDInt, err := strconv.ParseUint(snackID, 10, 32)
	if err != nil {
		return slot, snack, errors.New("Invalid request URL")
	}
	s, snackExists := GetSnack(snackIDInt)
	if !snackExists {
		return slot, snack, errors.New("Invalid snack ID")
	}
	snack = s

	var slotReq CoinSlot
	errDec := json.NewDecoder(r.Body).Decode(&slotReq)
	if errDec != nil {
		return slot, snack, errDec
	}
	
	slotResp, slotExists := GetCoinSlot(slotReq.Name)
	if !slotExists {
		return slot, snack, errors.New("Invalid slot ID")
	}

	if slotReq.Secret != slotResp.Secret {
		return slot, snack, errors.New("Invalid secret")
	}
	slot = slotResp
	return slot, snack, nil
}

func BuySnack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/snacks/:snack called")
	slot, snack, err := ValidateRequestBuySnack(r)
	if err != nil {
		fmt.Println("Error occured: ", err)
		returnError(w, http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		if slot.Balance > snack.Price {
			WithdrawCoinSlotBalance(slot, snack.Price)
		} else {
			fmt.Println("Not enough money to buy product")
			returnError(w, http.StatusBadRequest)
		}
	} else {
		returnError(w, http.StatusMethodNotAllowed)
	}
}

func main() {
	//FIXME it is better to use some ready URL router package, but I want to keep this sample as minimal as possible
    http.HandleFunc("/snacks", Snacks)
    http.HandleFunc("/snacks/", BuySnack)
    http.HandleFunc("/coin_slots", CoinSlots)
    http.HandleFunc("/coin_slots/", SingleCoinSlot)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Fatal error on ListenAndServe: ", err)
    }
}
