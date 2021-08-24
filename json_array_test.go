package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Edwin",
		Middlename: "Yordan",
		LastName:   "Laksono",
		Age:        30,
		Married:    false,
		Hobbies:    []string{"gaming", "reading", "coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Edwin","Middlename":"Yordan","LastName":"Laksono","Age":30,"Married":false,"Hobbies":["gaming","reading","coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesi",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesi",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","Middlename":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesi","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesi","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Middlename)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesi","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesi","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresess := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresess)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresess)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesi",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesi",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
