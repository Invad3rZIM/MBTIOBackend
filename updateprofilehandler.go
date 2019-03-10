package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//updates user info
func (h *Handler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {

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

	//update user values if they're part of the json string

	if name, ok := requestBody["name"].(string); ok {
		user.SetName(name)
	}

	if bio, ok := requestBody["bio"].(string); ok {
		user.SetBio(bio)
	}

	if sex, ok := requestBody["sex"].(string); ok {
		user.SetSex(sex)
	}

	if interest, ok := requestBody["interest"].(string); ok {
		user.SetInterest(interest)
	}

	if height, ok := requestBody["height"].(float64); ok {
		user.SetHeight(int(height))
	}

	if age, ok := requestBody["age"].(float64); ok {
		user.SetAge(int(age))
	}

	if mbti, ok := requestBody["mbti"].(string); ok {
		user.SetMBTI(mbti)
	}

	if lat, ok := requestBody["lat"].(float64); ok {
		user.SetLat(lat)
	}

	if long, ok := requestBody["long"].(float64); ok {
		user.SetLong(long)
	}

	if !user.ProfileReady {
		user.EvaluateProfileReady()
	}

	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}
