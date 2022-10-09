package main

type ConfirmWebhookResponse struct {
	Challenge string `json:"hub.challenge"`
}

type WebhookPayload struct {
	AspectType string `json:"aspect_type"`
	ObjectType string `json:"object_type"`
	ObjectID   string `json:"object_id"`
	OwnerID    string `json:"owner_id"`
}
