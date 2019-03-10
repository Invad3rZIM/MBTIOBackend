package main

import (
	"time"
)

type Message struct {
	Sid     int    `json:"sid"`
	Rid     int    `json:"rid"`
	Num     int    `json:"num"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

//text messages, autogenerates the timestamp
func NewMessage(sid int, rid int, num int, message string) *Message {
	return &Message{
		Sid:     sid,
		Rid:     rid,
		Num:     num,
		Message: message,
		Time:    FormatTime(time.Now().Format(time.RFC850)),
	}
}

//used for formatting timestamps later. dummy function for now
func FormatTime(t string) string {
	return t
}
