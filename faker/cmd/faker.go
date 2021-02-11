package main

import (
	"faker/internal/action"
	"log"
	"os"
	"time"
)

func main() {
	duration := action.GetDuration(os.Args)
	for {
		log.Println("Message sent.")
		time.Sleep(duration * time.Second)
	}
}
