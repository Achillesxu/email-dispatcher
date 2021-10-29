package dto

type KafkaResponse struct {
	To       string `json:"to,omitempty"`
	Subject  string `json:"subject,omitempty"`
	Template string `json:"template,omitempty"`
}
