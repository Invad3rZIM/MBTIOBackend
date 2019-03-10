package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//returns information about the user in question
func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "pin", "uid"); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//extract field values to variables for error checking/readability
	uid := int(requestBody["uid"].(float64))

	//extract field values to variables for error checking/readability
	pin := int(requestBody["pin"].(float64))

	//ensure pin is verified for user authentication
	if err := h.VerifyPin(uid, pin); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	user, err := h.GetUser(uid)

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}

//returns information about the user in question
func (h *Handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	all := []*User{}

	for _, u := range h.UserCache.cache {
		all = append(all, u)
	}

	json.NewEncoder(w).Encode(&all)
	w.WriteHeader(http.StatusOK)
}
