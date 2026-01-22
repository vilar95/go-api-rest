package response

import (
	"encoding/json"
	"go-api-rest/internal/dto"
	"net/http"
)

// JSON envia uma resposta JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Success envia uma resposta de sucesso
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	JSON(w, statusCode, data)
}

// Error envia uma resposta de erro
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, dto.ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	})
}

// ValidationError envia uma resposta de erro de validação
func ValidationError(w http.ResponseWriter, errors map[string]string) {
	JSON(w, http.StatusBadRequest, dto.ErrorResponse{
		Error:   "Erro de validação",
		Message: "Os dados fornecidos são inválidos",
		Details: errors,
	})
}

// Created envia uma resposta de recurso criado
func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

// NoContent envia uma resposta sem conteúdo
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
