package main

type CoinSlot struct {
    Name string
    Secret string
    Balance uint64
}

type Snack struct {
    ID      uint64
    Name    string
    Price   uint64 //in cents
}

type Coin struct {
    Coin uint64
}

var snacksTable = []Snack{
    {10, "Snickers", 220},
    {11, "Mars", 210},
}

var coinSlotMap = make(map[string]CoinSlot)

func GetSnacksTable() []Snack {
    return snacksTable
}

func CreateNewCoinSlot() CoinSlot {
	slot := CoinSlot{GetRandomName(), GetRandomSecret(), 0}
    coinSlotMap[slot.Name] = slot
    return slot
}

func UpdateCoinSlotBalance(cs CoinSlot, amount uint64) {
    cs.Balance += amount
    coinSlotMap[cs.Name] = cs
}

func GetCoinSlot(slotID string) (CoinSlot, bool) {
    slot, keyExists := coinSlotMap[slotID]
    return slot, keyExists
}
