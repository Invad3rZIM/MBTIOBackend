package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"bitbucket.org/ckvist/twilio/twirest"
)

type Twilio struct {
	//cache contains phone# -> twiliopin map
	cache      map[string]int
	accountSid string
	authToken  string
	from       string
	client     *twirest.TwilioClient
}

func NewTwilio() *Twilio {
	return &Twilio{
		cache:      make(map[string]int),
		accountSid: mustGetenv("TWILIO_ACCOUNT_SID"),
		authToken:  mustGetenv("TWILIO_AUTH_TOKEN"),
		from:       mustGetenv("TWILIO_NUMBER"),
		client:     twirest.NewClient(mustGetenv("TWILIO_ACCOUNT_SID"), mustGetenv("TWILIO_AUTH_TOKEN")),
	}
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

//adds user to cache, generating a 6 digit random pin as a string
func (t *Twilio) AddToCache(phone string) error {
	//gen pin
	twiliPin := GenTwilioPin(9999)

	if _, ok := t.cache[phone]; ok {
		return errors.New("error: phone number already in twilio cache")
	}

	err := t.TextPin(phone, twiliPin)

	if err != nil {
		return err
	}

	//add to cache
	t.cache[phone] = twiliPin

	return nil
}

//sends the pin to the phone number specified by the "to" parameter
func (t *Twilio) TextPin(to string, pin int) error {
	msg := twirest.SendMessage{
		Text: fmt.Sprintf("Your mbt.io verify pin is: %04d", pin),
		From: t.from,
		To:   to,
	}

	_, err := t.client.Request(msg)

	return err
}

func (t *Twilio) RemoveFromCache(phone string) {
	delete(t.cache, phone)
}

//checks the user is verified; or an error message explaining why not, nil if everything works as intended
func (t *Twilio) VerifyPhone(phone string, pin int) error {
	p, ok := t.cache[phone]

	if !ok {
		return errors.New("error: phone # not found")
	}

	if p != pin {
		return errors.New("error: invalid pin")
	}

	return nil
}

//filler id generator function. needs rework later
func GenTwilioPin(max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max)
}
