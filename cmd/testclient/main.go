package main

import (
	"context"
	"log"

	"github.com/kyleterry/tenyks/pkg/client"
)

func main() {
	c := client.NewClient("localhost:9999")

	log.Println(c.Connect(context.Background()))
}
