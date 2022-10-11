package main

type ConfirmWebhookResponse struct {
	Challenge string `json:"hub.challenge"`
}

type WebhookPayload struct {
	AspectType string `json:"aspect_type"`
	ObjectType string `json:"object_type"`
	ObjectID   int 	  `json:"object_id"`
	OwnerID    int    `json:"owner_id"`
}
