package handler

import (
	"encoding/json"
	"errors"
	"go-api-rest/internal/dto"
	"go-api-rest/internal/service"
	"go-api-rest/pkg/logger"
	"go-api-rest/pkg/response"
	customValidator "go-api-rest/pkg/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PersonalityHandler gerencia as requisições HTTP para personalidades
type PersonalityHandler struct {
	service service.PersonalityService
}

// NewPersonalityHandler cria uma nova instância do handler
func NewPersonalityHandler(service service.PersonalityService) *PersonalityHandler {
	return &PersonalityHandler{service: service}
}

// Home retorna a página inicial da API
func (h *PersonalityHandler) Home(w http.ResponseWriter, r *http.Request) {
	response.Success(w, http.StatusOK, dto.SuccessResponse{
		Message: "Bem-vindo à API REST de Personalidades em Go!",
	})
}

// GetAll retorna todas as personalidades
func (h *PersonalityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	personalities, err := h.service.GetAll()
	if err != nil {
		logger.Errorf("Erro ao buscar personalidades: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao buscar personalidades")
		return
	}

	response.Success(w, http.StatusOK, personalities)
}

// GetByID retorna uma personalidade por ID
func (h *PersonalityHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	personality, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrPersonalityNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		logger.Errorf("Erro ao buscar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao buscar personalidade")
		return
	}

	response.Success(w, http.StatusOK, personality)
}

// Create cria uma nova personalidade
func (h *PersonalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePersonalityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	// Validar dados
	if validationErrors := customValidator.ValidateStruct(req); validationErrors != nil {
		response.ValidationError(w, validationErrors)
		return
	}

	personality, err := h.service.Create(&req)
	if err != nil {
		if errors.Is(err, service.ErrPersonalityAlreadyExists) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		logger.Errorf("Erro ao criar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao criar personalidade")
		return
	}

	response.Created(w, personality)
}

// Update atualiza uma personalidade existente
func (h *PersonalityHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var req dto.UpdatePersonalityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Dados inválidos")
		return
	}

	// Validar dados
	if validationErrors := customValidator.ValidateStruct(req); validationErrors != nil {
		response.ValidationError(w, validationErrors)
		return
	}

	personality, err := h.service.Update(uint(id), &req)
	if err != nil {
		if errors.Is(err, service.ErrPersonalityNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, service.ErrPersonalityAlreadyExists) {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}
		logger.Errorf("Erro ao atualizar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao atualizar personalidade")
		return
	}

	response.Success(w, http.StatusOK, personality)
}

// Delete remove uma personalidade
func (h *PersonalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if errors.Is(err, service.ErrPersonalityNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		logger.Errorf("Erro ao deletar personalidade: %v", err)
		response.Error(w, http.StatusInternalServerError, "Erro ao deletar personalidade")
		return
	}

	response.NoContent(w)
}
