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

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	fmt.Printf("%+v", config)

}
