package main

import (
	"log"

	"github.com/eminetto/openapi2test/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
