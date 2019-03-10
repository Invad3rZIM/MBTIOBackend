package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "pin", "sid", "rid", "num", "message"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for error checking/readability
	sid := int(requestBody["sid"].(float64))
	pin := int(requestBody["pin"].(float64))

	//ensure pin is verified for user authentication
	if err := h.VerifyPin(sid, pin); err != nil {
		fmt.Fprintln(w, err.Error()+" e")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rid := int(requestBody["rid"].(float64))
	num := int(requestBody["num"].(float64))
	message := requestBody["message"].(string)

	m := NewMessage(sid, rid, num, message)

	//add message to cache
	h.PostMessage(m)

	json.NewEncoder(w).Encode(&m)
	w.WriteHeader(http.StatusOK)
}
