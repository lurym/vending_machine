package main

type Snack struct {
    Id      uint64
    Name    string
    Price   uint64 //in cents
}

var snacksTable []Snack = []Snack{
    {10, "Snickers", 220},
    {11, "Mars", 210},
}
