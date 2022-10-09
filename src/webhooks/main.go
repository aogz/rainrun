package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", ReceiveWebhook)
	http.HandleFunc("/webhook/confirm", ConfirmWebhook)
	http.ListenAndServe(":3001", nil)
}
