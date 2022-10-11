package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func CreateWebhook() error {
	callbackURL := fmt.Sprintf("https://%s/webhook/", os.Getenv("APP_DOMAIN"))
	logger.info.Printf("Setting callback url to %s\n", callbackURL)
	queryParams := url.Values{
		"client_id":     {os.Getenv("STRAVA_CLIENT_ID")},
		"client_secret": {os.Getenv("STRAVA_CLIENT_SECRET")},
		"verify_token":  {"STRAVA_VERIFY_TOKEN"},
		"callback_url":  {callbackURL},
	}

	createWebhookURL := url.URL{
		Scheme:   "https",
		Host:     "www.strava.com",
		Path:     "/api/v3/push_subscriptions",
		RawQuery: queryParams.Encode(),
	}

	response, err := http.Post(createWebhookURL.String(), "application/x-www-form-urlencoded", nil)
	if err != nil {
		return fmt.Errorf("error sending post request: %s", err)
	}

	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		logger.err.Fatalln(err)
	} else {
		logger.info.Println(string(responseBytes))
	}
	return nil
}

func ConfirmWebhook(w http.ResponseWriter, r *http.Request) {
	data := ConfirmWebhookResponse{
		Challenge: r.URL.Query().Get("hub.challenge"),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ReceiveWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ConfirmWebhook(w, r)
	} else {
		var event WebhookPayload
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			errorMessage := fmt.Sprintf("unable to parse: %s", err.Error())
			logger.err.Fatalln(errorMessage)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		logger.info.Printf(
			"Event received: event_type=%s, object_type=%s, id=%d, owner_id=%d\n",
			event.AspectType,
			event.ObjectType,
			event.ObjectID,
			event.OwnerID,
		)

		fmt.Fprintf(w, "Person: %+v", event)
	}
}
