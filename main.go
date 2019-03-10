package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine library
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	//	fmt.Fprintln(w, "Rob and Kirk Test!")

	var requestBody map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println(err.Error())
		return
	}

	name := requestBody["name"].(string)

	fmt.Fprintln(w, name)
}

func main() {

	h := NewHandler()

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/matches/get", h.GetMatchesHandler) //need to do

	http.HandleFunc("/messages/get", h.GetMessagesHandler)
	http.HandleFunc("/messages/send", h.SendMessageHandler)

	http.HandleFunc("/twilio/send", h.SendTwilioHandler)
	http.HandleFunc("/twilio/verify", h.VerifyTwilioHandler)

	http.HandleFunc("/profile/update", h.UpdateProfileHandler)
	http.HandleFunc("/profile/ready", h.CheckProfileReadyHandler)
	http.HandleFunc("/profile/self", h.GetUserHandler)
	http.HandleFunc("/profile/all", h.GetAllUsersHandler)

	appengine.Main() // Starts the server to receive requests
}

/*
	var requestBody map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		fmt.Println(err.Error())
		//	WriteBadRequestErrorResponse(&w)
		return
	}

	this is how i type my words like this i'm unsure what

	//conver to ints here for
	name := requestBody["name"].(string)

	user := users.NewUser(name, 0, 3, 0, 0, 0)
	h.UserCache.Users[user.UserID] = user

	//add user to database
	h.DB.AddUser(user)
	h.DB.AddHeart(heart.NewHeart(user))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//return copy of user for client records
	json.NewEncoder(w).Encode(&user)
*/
