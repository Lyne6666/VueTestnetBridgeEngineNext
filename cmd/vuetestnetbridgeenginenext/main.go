// cmd/vuetestnetbridgeenginenext/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"vuetestnetbridgeenginenext/internal/vuetestnetbridgeenginenext"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := vuetestnetbridgeenginenext.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
