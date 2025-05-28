package main

import (
	"fmt"
	"log"
	"github.com/yourorg/syniac-sdk-go/syniac"
)

func main() {
	client := syniac.NewClient("your_api_key")

	instances, err := client.ListInstances()
	if err != nil {
		log.Fatal(err)
	}

	for _, inst := range instances {
		fmt.Println(inst.ID, inst.Name, inst.Status)
	}
}
