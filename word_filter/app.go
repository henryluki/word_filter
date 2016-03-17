package main

import (
	"log"
	"net/http"
)

func main() {
	// load dict and watch file change
	InitLoader()

	log.Println("Run word_filter app on", 8001)
	// routes
	http.HandleFunc("/word/is_valid", VerifyWordsHandler)
	// run server
	http.ListenAndServe(":8001", nil)
}
