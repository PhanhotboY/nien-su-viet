package utils

type OperationResponse struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func NewOperationResponse(id string, success bool, message string) *OperationResponse {
	return &OperationResponse{
		ID:      id,
		Success: success,
		Message: message,
	}
}
