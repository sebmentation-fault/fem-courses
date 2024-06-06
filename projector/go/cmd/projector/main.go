package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sebmentation-fault/projector/pkg/projector"
)


func main() {

	opts, err := projector.GetOpts()

	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	config, err := projector.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	proj := projector.NewProjector(config)

	if config.Operation == projector.Print {
		fmt.Println("printing...")
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			fmt.Printf("data: %+v\n", data)
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("this line should never be reached %v", err)
			}
			fmt.Printf("%v\n", string(jsonString))
		} else if value, ok := proj.GetValue(config.Args[0]); ok {
			fmt.Printf("data: %+v\n", value)
			fmt.Printf("%v\n", value)
		}
	}

	if config.Operation == projector.Add {
		fmt.Println("adding...")
		proj.SetValue(config.Args[0], config.Args[1])
		proj.Save()
	}

	if config.Operation == projector.Remove {
		fmt.Println("removing...")
		proj.RemoveValue(config.Args[0])
		proj.Save()
	}
}

