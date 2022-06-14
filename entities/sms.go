package entities

type KaleyraSMSRequest struct {
	Message          string   `json:"message" binding:"required"`
	From             string   `json:"from" binding:"required"`
	To               string   `json:"to" binding:"required"`
	Webhook          string   `json:"webhook"`
	TemplateID       string   `json:"template_id"`
	Schedule         string   `json:"schedule"`
	AdditionalParams []string `json:"additional_params"`
}

type SMSResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"service_response"`
}
