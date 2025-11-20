package main

import (
	"log"

	"thunderbird.zap/idp/internal/configuration"
)

func main() {
	app := &application{
		config: configuration.Init(),
	}
	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal("Unable to Start Server")
	}
}
