package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) VerifyTwilioHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "phone", "twiliopin"); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	phone := requestBody["phone"].(string)
	twilioPin := int(requestBody["twiliopin"].(float64))

	//verify phone number in cache
	if err := h.Twilio.VerifyPhone(phone, twilioPin); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//adds user to cache
	u := h.UserCache.CacheUserFromPhone(phone)

	//removes phone from cache
	h.Twilio.RemoveFromCache(phone)

	json.NewEncoder(w).Encode(&u)
	w.WriteHeader(http.StatusOK)
}
