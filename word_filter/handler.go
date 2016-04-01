package main

import (
	"log"
	"net/http"
)

// handle funcs
func VerifyWordsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		text := req.FormValue("v")
		var hit bool
		var review_level int

		level := VerifyWords(text)
		if level == 0 {
			hit = false
			review_level = 0
		} else {
			hit = true
			// if level equals to 2, need to predict text label
			if level == 2 {
				label := PredictText(text)
				if label == 0 {
					review_level = 3 // need ban
				} else {
					review_level = level
				}
			} else {
				review_level = level
			}
		}
		log.Println("[INFO] text:", text, "hit:", hit, "level:", level)
		res := HitResponse{Hit: hit, Level: review_level}
		w.Write(RenderJson(res))
	}
}
