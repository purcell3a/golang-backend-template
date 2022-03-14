package main

import (
	"fmt"
	"log"

)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("database ensured!")
}
