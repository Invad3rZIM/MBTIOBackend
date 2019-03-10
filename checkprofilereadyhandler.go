package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//checks to see if all requisite information is filled out for the user to begin participating in the dating pool
func (h *Handler) CheckProfileReadyHandler(w http.ResponseWriter, r *http.Request) {

	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "uid", "pin"); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for error checking/readability
	uid := int(requestBody["uid"].(float64))
	pin := int(requestBody["pin"].(float64))

	//ensure pin is verified for user authentication
	if err := h.VerifyPin(uid, pin); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//if user if found and verified
	user, err := h.GetUser(uid)

	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user.EvaluateProfileReady()

	json.NewEncoder(w).Encode(&user.ProfileReady)
	w.WriteHeader(http.StatusOK)
}
