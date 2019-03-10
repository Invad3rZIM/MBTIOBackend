package main

import (
	"errors"
	"fmt"
)

type Handler struct {
	*Twilio
	*UserCache
	*Database
	*MessageCache
	*MBTIChart
}

func NewHandler() *Handler {
	/*db := NewDatabase()

	if we can't load the cloud database, all is lost... TRIPLE CHECK THIS ALTER
	if err != nil {
		log.Fatal(err.Error())
	}*/

	h := Handler{
		NewTwilio(),
		NewUserCache(),
		NewDatabase(),
		NewMessageCache(),
		NewMBTIChart(),
	}

	(&h).CreateTestUsers(100)

	return &h
}

//authenticates the user via numeric pin... replace with JWT if time permits later
func (h *Handler) VerifyPin(uid int, pin int) error {
	u, err := h.GetUser(uid)

	//if user cannot be found for whatever reason
	if err != nil {
		return err
	}

	if !u.CheckPin(pin) {
		return errors.New("error: invalid pin")
	}

	return nil
}

//VerifyBody is a helper function to ensure all http requests contain the requisite fields returns error if fields missing
func (h *Handler) VerifyBody(body map[string]interface{}, str ...string) error {
	for _, s := range str {
		fmt.Println(s)
		if _, ok := body[s]; !ok {
			return errors.New("error: missing field: " + s)
		}
	}

	return nil
}
