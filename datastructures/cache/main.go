package main

import (
	"fmt"
)

func main() {
	uc := UsersCache{
		Limit: 100,
	}
	err := uc.Add("Pepper")
	if err != nil {
	}
	val, err := uc.Find("Pepper")
	fmt.Println(val)
}
