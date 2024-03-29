package main

type Athlete struct {
	ID int `json:"id"`
}

type TokenResponse struct {
	AccessToken  string  `json:"access_token"`
	ExpiresAt    int     `json:"expires_at"`
	ExpiresIn    int     `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	Athlete      Athlete `json:"athlete"`
}
