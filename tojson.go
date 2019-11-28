package main

import (
	"encoding/json"
	"fmt"
)

func toJSON(langStats map[string]int) {
	json, err := json.Marshal(langStats)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
		return
	}

	fmt.Println(string(json))
}
