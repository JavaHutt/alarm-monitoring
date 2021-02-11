package main

import (
	"encoding/json"
	"faker/internal/action"
	"fmt"
	"log"
	"os"
)

func main() {
	duration := action.GetDuration(os.Args)
	fmt.Println("duration", duration)
	alarm := action.GetRandomAlarm()
	json, err := json.Marshal(alarm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))

}
