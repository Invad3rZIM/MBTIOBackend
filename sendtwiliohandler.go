package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Twilio Send TextMessage
func (h *Handler) SendTwilioHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "phone"); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	phone := requestBody["phone"].(string)

	err := h.Twilio.AddToCache(phone)

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
