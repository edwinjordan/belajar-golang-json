package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Edwin",
		Middlename: "Yordan",
		LastName:   "Laksono",
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)
}
