package main

import (
	"fmt"
	"log"

	"github.com/FozzyHosting/go-winvps"
)

func pagination() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	rOpts := &winvps.RequestOptions{Limit: 10, Page: 1}

	for {
		machines, page, err := winClient.GetMachines(rOpts)
		if err != nil {
			fmt.Println(err)
		}
		for _, m := range machines {
			fmt.Println(m.Name, m.Status)
		}
		if page.NextPage() == 0 {
			break
		}
		rOpts.Page = page.NextPage()
	}
}
