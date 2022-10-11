package main

import (
	"log"
	"net/http"
	"os"
)

var logger Logger
type Logger struct {
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
}

func init() {
	logger.info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    logger.warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    logger.err = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// Create application webhook
	go CreateWebhook()

	http.HandleFunc("/webhook/", ReceiveWebhook)
	http.HandleFunc("/webhook/confirm/", ConfirmWebhook)
	http.ListenAndServe(":3001", nil)
}
