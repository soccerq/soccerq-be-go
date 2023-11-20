package main

import (
	"fmt"
	"soccerq/internal/app"
)

func main() {
	config := app.Config{
		Name:     "SoccerQ",
		Version:  "0.0.1",
		Port:     ":3000",
		DB_READ:  "",
		DB_WRITE: "",
	}

	fmt.Printf("%+v", config)
}
