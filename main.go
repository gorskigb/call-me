package main

import (
	"encoding/json"
	"net/http"
	"call-me/contract"
)

func main() {
	//TODO: save request in database and handle with separate routine
	http.HandleFunc("/call/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed. ", http.StatusMethodNotAllowed)
			return
		}

		var callBody contract.GetCall

		//TODO: handle excpetion
		json.NewDecoder(r.Body).Decode(&callBody)

		http.Get(callBody.Url)
	})

	//TODO Address should be configurable
	http.ListenAndServe(":8181", nil)
}
