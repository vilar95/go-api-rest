package dto

// CreatePersonalityRequest representa os dados para criar uma personalidade
type CreatePersonalityRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=100"`
	History string `json:"history" validate:"required,min=10,max=5000"`
}

// UpdatePersonalityRequest representa os dados para atualizar uma personalidade
type UpdatePersonalityRequest struct {
	Name    string `json:"name" validate:"omitempty,min=3,max=100"`
	History string `json:"history" validate:"omitempty,min=10,max=5000"`
}

// PersonalityResponse representa a resposta da API
type PersonalityResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	History string `json:"history"`
}

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

// SuccessResponse representa uma resposta de sucesso genérica
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
