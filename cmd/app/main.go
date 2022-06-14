package main

import (
	"log"
	"mux-api-boilerplate/init/application"
)

func main() {
	err := application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
