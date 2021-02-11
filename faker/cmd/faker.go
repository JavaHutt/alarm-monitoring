package main

import (
	"encoding/json"
	"faker/internal/action"
	"fmt"
	"log"
)

func main() {
	alarm := action.GetRandomAlarm()
	json, err := json.Marshal(alarm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))

}
