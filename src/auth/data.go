package main

type Athlete struct {
}

type TokenResponse struct {
	AccessToken  string  `json:"access_token"`
	ExpiresAt    string  `json:"expires_at"`
	ExpiresIn    int     `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	Athlete      Athlete `json:"athlete"`
}

type ConfirmWebhookResponse struct {
	Challenge string `json:"hub.challenge"`
}

type WebhookPayload struct {
	AspectType string `json:"aspect_type"`
	ObjectType string `json:"object_type"`
	ObjectID   string `json:"object_id"`
	OwnerID    string `json:"owner_id"`
}
