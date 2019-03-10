package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetMatchesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "pin", "uid"); err != nil {
		return
	}

	//extract field values to variables for error checking/readability
	uid := int(requestBody["uid"].(float64))
	//extract field values to variables for error checking/readability
	pin := int(requestBody["pin"].(float64))

	//ensure pin is verified for user authentication
	if err := h.VerifyPin(uid, pin); err != nil {
		return
	}

	max := 10 //optional messages max, defaulted to 10

	if m, ok := requestBody["max"].(float64); ok {
		max = int(m)
	}

	matches, err := h.GetMatches(uid, max)

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	ww := Wrapper{M: matches}

	json.NewEncoder(w).Encode(&ww)

	w.WriteHeader(http.StatusOK)
}

type Wrapper struct {
	M []*User
}
