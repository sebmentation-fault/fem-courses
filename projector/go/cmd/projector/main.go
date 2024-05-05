package main

import (
	"fmt"
	"log"

	"github.com/sebmentation-fault/projector/pkg/projector"
)


func main() {

	opts, err := projector.GetOpts()

	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	fmt.Printf("%+v", opts)

}
