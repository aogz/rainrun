package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func index(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "templates/index.html")
}

func RedirectToOAuth(w http.ResponseWriter, r *http.Request) {
	redirectURI := fmt.Sprintf("http://%s/oauth/callback", os.Getenv("APP_DOMAIN"))
	queryParams := url.Values{
		"client_id":     {os.Getenv("STRAVA_CLIENT_ID")},
		"redirect_uri":  {redirectURI},
		"response_type": {"code"},
		"scope":         {"profile:read_all,activity:read_all,activity:write"},
	}

	authorizeURL := url.URL{
		Scheme:   "https",
		Host:     "www.strava.com",
		Path:     "/oauth/authorize",
		RawQuery: queryParams.Encode(),
	}

	http.Redirect(w, r, authorizeURL.String(), http.StatusTemporaryRedirect)
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "Error parsing URL: %s", err)
		return
	}

	code := url.Query().Get("code")
	if code == "" {
		fmt.Fprintf(w, "Error: code is empty")
		return
	}

	token, err := OAauthComplete(code)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.Write([]byte(fmt.Sprintf("Token: %s", token.AccessToken)))
}

func OAauthComplete(code string) (TokenResponse, error) {
	var token TokenResponse
	queryParams := url.Values{
		"client_id":     {os.Getenv("STRAVA_CLIENT_ID")},
		"client_secret": {os.Getenv("STRAVA_CLIENT_SECRET")},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	tokenURL := url.URL{
		Scheme:   "https",
		Host:     "www.strava.com",
		Path:     "/oauth/token",
		RawQuery: queryParams.Encode(),
	}
	response, err := http.Post(tokenURL.String(), "application/x-www-form-urlencoded", nil)
	if err != nil {
		return token, fmt.Errorf("error sending post request: %s", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return token, fmt.Errorf("error reading response body: %s", err)
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return token, fmt.Errorf("error unmarhalling response: %s", err)
	}

	return token, nil
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", RedirectToOAuth)
	http.HandleFunc("/oauth/callback", OAuthCallback)
	http.ListenAndServe(":3000", nil)
}
