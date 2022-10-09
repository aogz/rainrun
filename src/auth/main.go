package main

import (
	"net/http"
)

func index(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/index.html")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", RedirectToOAuth)
	http.HandleFunc("/oauth/callback", OAuthCallback)
	http.ListenAndServe(":3000", nil)
}
