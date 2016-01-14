package main

import (
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
			review_level = level
		}
		res := Response{Hit: hit, Level: review_level}
		w.Write(ToJson(res))
	}
}
