package dto

type KafkaResponse struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}
