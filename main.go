package main

import (
	"github.com/aerostatka/db-testing/internal/app"
	"os"
)

func main()  {
	if len(os.Args) < 2 {
		panic("Call should have at least 1 argument")
	}

	mode := os.Args[1]
	app.Start(mode)
}
