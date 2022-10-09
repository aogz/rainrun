package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func CreateWebhook() error {
	callbackURL := fmt.Sprintf("http://%s/webhook/confirm", os.Getenv("APP_DOMAIN"))
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
	return nil
}

func ConfirmWebhook(w http.ResponseWriter, r *http.Request) {
	data := ConfirmWebhookResponse{
		Challenge: "foobar",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func ReceiveWebhook(w http.ResponseWriter, r *http.Request) {
	var event WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf(
		"Event received: event_type=%s, object_type=%s, id=%s, owner_id=%s\n",
		event.AspectType,
		event.ObjectType,
		event.ObjectID,
		event.OwnerID,
	)
	fmt.Fprintf(w, "Person: %+v", event)
}
