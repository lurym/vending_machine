package main

type CoinSlot struct {
    Name string
    Secret string
    Balance int64
}

type Snack struct {
    ID      uint64
    Name    string
    Price   uint64 //in cents
}

var snacksTable = []Snack{
    {10, "Snickers", 220},
    {11, "Mars", 210},
}

var coinSlotTable = []CoinSlot{}

func CreateNewCoinSlot() CoinSlot {
	slot := CoinSlot{GetRandomName(), GetRandomSecret(), 0}
	coinSlotTable = append(coinSlotTable, slot) //FIXME this is not thread safe
    return slot
}
