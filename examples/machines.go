package main

import (
	"fmt"
	"log"

	"github.com/FozzyHosting/go-winvps"
)

func getMachines() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// short form
	machines, _, err := winClient.GetMachines()

	// complete form
	machines, _, err = winClient.GetMachinesFull()

	for _, m := range machines {
		fmt.Println(m.Name, m.Status)
	}
}

func createMachine() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	createOpts := &winvps.CreateMachineOptions{
		ProductID:  1,
		TemplateID: 2,
		LocationID: 3,
	}

	machineName, jobs, err := winClient.CreateMachine(createOpts)
	if err != nil {
		log.Fatalf("Failed to create machine: %v", err)
	}
	fmt.Printf("Machine %s create accepted, job id: %s", machineName, jobs[0].ID)
}

func updateMachine() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	updateOpts := &winvps.UpdateMachineOptions{
		Password: "newpassword",
		AddDisk:  10,
	}
	machineName := "VPS0123"
	jobs, err := winClient.UpdateMachine(machineName, updateOpts)
	if err != nil {
		log.Fatalf("Failed to update machine: %v", err)
	}
	fmt.Printf("Machine %s update accepted, job id: %s", machineName, jobs[0].ID)
}

func reinstallMachine() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	reinstallOpts := &winvps.ReinstallMachineOptions{
		Password: "newpassword",
	}
	machineName := "VPS0123"
	jobs, err := winClient.ReinstallMachine(machineName, reinstallOpts)
	if err != nil {
		log.Fatalf("Failed to reintall machine: %v", err)
	}
	fmt.Printf("Machine %s reinstall accepted, job id: %s", machineName, jobs[0].ID)
}

func deleteMachine() {
	winClient, err := winvps.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	machineName := "VPS0123"
	jobs, err := winClient.DeleteMachine(machineName)
	if err != nil {
		log.Fatalf("Failed to delete machine: %v", err)
	}
	fmt.Printf("Machine %s delete accepted, job id: %s", machineName, jobs[0].ID)
}