package main

import (
	"fmt"
	"log"

	"github.com/FozzyHosting/go-winvps"
)

func GetProducts() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	products, _, err := winClient.GetProducts()
	for _, p := range products {
		fmt.Println(p.ID, p.Name)
	}
}
