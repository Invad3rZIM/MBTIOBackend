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

func (h *Handler) GetFakeMessagesHandler(w http.ResponseWriter, r *http.Request) {
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

	//dummy message array in place of actually getting new messages
	m := []*Message{NewMessage(0, 1, 0, "hey"),
		NewMessage(0, 1, 1, "hey hows it going??"),
		NewMessage(0, 0, 2, "pretty well, what about you???"),
		NewMessage(0, 1, 4, "i'm doing alright. i just broke up with my boyfriend and looking for someone who has a compatible mbti personallity to do me dirty :)"),
		NewMessage(1, 0, 5, "I'm kirk. I'm an ENTJ!"),
		NewMessage(0, 1, 6, "I heart ENTJs are kinda psycho... are you kinda psycho??"),
		NewMessage(1, 0, 8, "Kinda"),
		NewMessage(0, 1, 9, "That's okay, i'm kinda psycho too!"),
		NewMessage(1, 0, 10, "Something good can work.. #?"),
		NewMessage(0, 1, 11, "7323201127, My names Sasha!")}

	json.NewEncoder(w).Encode(&m)
}
