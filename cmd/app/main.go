package main

import (
	"log"
	"mux-api-boilerplate/init"
)

func main() {
	err := init.Api()
	if err != nil {
		log.Fatal(err)
	}
}
