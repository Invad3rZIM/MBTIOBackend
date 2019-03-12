package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "pin", "sid"); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//extract field values to variables for error checking/readability
	sid := int(requestBody["sid"].(float64))

	//extract field values to variables for error checking/readability
	pin := int(requestBody["pin"].(float64))

	//ensure pin is verified for user authentication
	if err := h.VerifyPin(sid, pin); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	m := h.GetMessages(sid)

	json.NewEncoder(w).Encode(&m)
}
