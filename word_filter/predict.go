package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	PREDICT_HOST = "http://127.0.0.1:8006/classify"
)

func PredictText(text string) string {
	resp, err := http.PostForm(PREDICT_HOST,
		url.Values{"text": {text}})
	if err != nil {
		// handle error
		log.Println("[ERROR]: request for Predict Host error!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Println("[ERROR]: response from Predict Host error!")
	}
	d := DecodeJson(body)
	log.Println("[INFO] label:" + d.Label + "text:" + d.Text)
	log.Println(d)
	return d.Label
}
