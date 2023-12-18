package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Packs struct {
	PackSizes []int `json:"packSizes"`
}

func getPackSizes() (Packs, error) {
	file, err := os.Open("./packSizes.json")

	if err != nil {
		return Packs{}, fmt.Errorf("Error opening file:", err)
	}

	defer file.Close()

	var packs Packs

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&packs)

	if err != nil {
		return Packs{}, fmt.Errorf("Error decoding JSON:", err)
	}

	return packs, nil
}
