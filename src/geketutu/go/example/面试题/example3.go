package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	t := map[string]string{
		"name": "Bob",
		"age":  "18",
	}

	jsonStr, _ := json.Marshal(t)
	fmt.Println(jsonStr)
	UnjsonStr := json.Unmarshal(jsonStr, &t)
	fmt.Println(UnjsonStr)
}
