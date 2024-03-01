package main

import (
	"log"
	"os"

	"github.com/merliot/skeleton"
)

//go:generate go run main.go
func main() {
	skeleton := skeleton.New("proto", "skeleton", "proto").(*skeleton.Skeleton)
	if err := skeleton.GenerateUf2s("../.."); err != nil {
		log.Println("Error generating UF2s:", err)
		os.Exit(1)
	}
}
