package main

import (
	"encoding/json"
	"log"
)

type Response struct {
	Hit   bool `json:"hit"`
	Level int  `json:"level"`
}

func ToJson(res Response) []byte {
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println("Json decode error.")
	}
	return jsonData
}
