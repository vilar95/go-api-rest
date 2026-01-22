package service

import (
	"errors"
	"go-api-rest/internal/dto"
	"go-api-rest/internal/repository"
	"go-api-rest/models"

	"gorm.io/gorm"
)

var (
	ErrPersonalityNotFound      = errors.New("personalidade não encontrada")
	ErrPersonalityAlreadyExists = errors.New("já existe uma personalidade com esse nome")
	ErrInvalidID                = errors.New("ID inválido")
)

// PersonalityService define a interface para lógica de negócio
type PersonalityService interface {
	Create(req *dto.CreatePersonalityRequest) (*dto.PersonalityResponse, error)
	GetAll() ([]dto.PersonalityResponse, error)
	GetByID(id uint) (*dto.PersonalityResponse, error)
	Update(id uint, req *dto.UpdatePersonalityRequest) (*dto.PersonalityResponse, error)
	Delete(id uint) error
}

type personalityService struct {
	repo repository.PersonalityRepository
}

// NewPersonalityService cria uma nova instância do serviço
func NewPersonalityService(repo repository.PersonalityRepository) PersonalityService {
	return &personalityService{repo: repo}
}

func (s *personalityService) Create(req *dto.CreatePersonalityRequest) (*dto.PersonalityResponse, error) {
	// Verificar se já existe uma personalidade com esse nome
	exists, err := s.repo.ExistsByName(req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPersonalityAlreadyExists
	}

	personality := &models.Personality{
		Name:    req.Name,
		History: req.History,
	}

	if err := s.repo.Create(personality); err != nil {
		return nil, err
	}

	return s.toDTO(personality), nil
}

func (s *personalityService) GetAll() ([]dto.PersonalityResponse, error) {
	personalities, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.PersonalityResponse, len(personalities))
	for i, p := range personalities {
		response[i] = *s.toDTO(&p)
	}

	return response, nil
}

func (s *personalityService) GetByID(id uint) (*dto.PersonalityResponse, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}

	personality, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPersonalityNotFound
		}
		return nil, err
	}

	return s.toDTO(personality), nil
}

func (s *personalityService) Update(id uint, req *dto.UpdatePersonalityRequest) (*dto.PersonalityResponse, error) {
	if id == 0 {
		return nil, ErrInvalidID
	}

	personality, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPersonalityNotFound
		}
		return nil, err
	}

	// Atualizar apenas campos não vazios
	if req.Name != "" {
		// Verificar se o novo nome já existe em outra personalidade
		exists, err := s.repo.ExistsByName(req.Name)
		if err != nil {
			return nil, err
		}
		if exists && personality.Name != req.Name {
			return nil, ErrPersonalityAlreadyExists
		}
		personality.Name = req.Name
	}

	if req.History != "" {
		personality.History = req.History
	}

	if err := s.repo.Update(personality); err != nil {
		return nil, err
	}

	return s.toDTO(personality), nil
}

func (s *personalityService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}

	// Verificar se existe antes de deletar
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPersonalityNotFound
		}
		return err
	}

	return s.repo.Delete(id)
}

// toDTO converte o modelo para DTO
func (s *personalityService) toDTO(p *models.Personality) *dto.PersonalityResponse {
	return &dto.PersonalityResponse{
		ID:      p.ID,
		Name:    p.Name,
		History: p.History,
	}
}
