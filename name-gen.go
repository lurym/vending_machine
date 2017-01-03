package main
//Custom name generator similar to docker container name generator

import (
    "math/rand"
    "fmt"
)

var left = []string { 
    "admiring",
    "adoring",
    "affectionate",
    "agitated",
    "amazing",
    "angry",
    "awesome",
}

var right = []string {
    "wilson",
    "wing",
    "wozniak",
    "wright",
    "yalow",
    "yonath",
}

var singleSecret = "secret_password"

func GetRandomName() string {
    return fmt.Sprintf("%s_%s", left[rand.Intn(len(left))], right[rand.Intn(len(right))])
}

func GetRandomSecret() string {
    //this is not a secret FIXME
    //should generate secure random passwords
    return singleSecret
}
