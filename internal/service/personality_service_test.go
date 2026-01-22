package service

import (
	"errors"
	"go-api-rest/internal/dto"
	"go-api-rest/models"
	"testing"

	"gorm.io/gorm"
)

// Mock do repository para testes
type mockPersonalityRepository struct {
	personalities map[uint]*models.Personality
	nextID        uint
}

func newMockRepository() *mockPersonalityRepository {
	return &mockPersonalityRepository{
		personalities: make(map[uint]*models.Personality),
		nextID:        1,
	}
}

func (m *mockPersonalityRepository) Create(personality *models.Personality) error {
	personality.ID = m.nextID
	m.personalities[m.nextID] = personality
	m.nextID++
	return nil
}

func (m *mockPersonalityRepository) FindAll() ([]models.Personality, error) {
	personalities := make([]models.Personality, 0, len(m.personalities))
	for _, p := range m.personalities {
		personalities = append(personalities, *p)
	}
	return personalities, nil
}

func (m *mockPersonalityRepository) FindByID(id uint) (*models.Personality, error) {
	p, exists := m.personalities[id]
	if !exists {
		return nil, gorm.ErrRecordNotFound
	}
	return p, nil
}

func (m *mockPersonalityRepository) Update(personality *models.Personality) error {
	if _, exists := m.personalities[personality.ID]; !exists {
		return gorm.ErrRecordNotFound
	}
	m.personalities[personality.ID] = personality
	return nil
}

func (m *mockPersonalityRepository) Delete(id uint) error {
	if _, exists := m.personalities[id]; !exists {
		return gorm.ErrRecordNotFound
	}
	delete(m.personalities, id)
	return nil
}

func (m *mockPersonalityRepository) ExistsByName(name string) (bool, error) {
	for _, p := range m.personalities {
		if p.Name == name {
			return true, nil
		}
	}
	return false, nil
}

// Testes
func TestCreatePersonality_Success(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	req := &dto.CreatePersonalityRequest{
		Name:    "Alan Turing",
		History: "Matemático e cientista da computação britânico",
	}

	result, err := service.Create(req)

	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	if result.Name != req.Name {
		t.Errorf("Esperava nome %s, mas obteve %s", req.Name, result.Name)
	}

	if result.ID == 0 {
		t.Error("ID não foi gerado")
	}
}

func TestCreatePersonality_DuplicateName(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	req := &dto.CreatePersonalityRequest{
		Name:    "Alan Turing",
		History: "Matemático e cientista da computação britânico",
	}

	// Primeira criação deve funcionar
	_, err := service.Create(req)
	if err != nil {
		t.Fatalf("Primeira criação falhou: %v", err)
	}

	// Segunda criação com mesmo nome deve falhar
	_, err = service.Create(req)
	if err == nil {
		t.Error("Esperava erro de nome duplicado, mas não obteve erro")
	}

	if !errors.Is(err, ErrPersonalityAlreadyExists) {
		t.Errorf("Esperava ErrPersonalityAlreadyExists, mas obteve: %v", err)
	}
}

func TestGetByID_Success(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	// Criar uma personalidade
	req := &dto.CreatePersonalityRequest{
		Name:    "Alan Turing",
		History: "Matemático e cientista da computação britânico",
	}
	created, _ := service.Create(req)

	// Buscar por ID
	result, err := service.GetByID(created.ID)

	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	if result.ID != created.ID {
		t.Errorf("Esperava ID %d, mas obteve %d", created.ID, result.ID)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	_, err := service.GetByID(999)

	if err == nil {
		t.Error("Esperava erro, mas não obteve erro")
	}

	if !errors.Is(err, ErrPersonalityNotFound) {
		t.Errorf("Esperava ErrPersonalityNotFound, mas obteve: %v", err)
	}
}

func TestGetByID_InvalidID(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	_, err := service.GetByID(0)

	if err == nil {
		t.Error("Esperava erro, mas não obteve erro")
	}

	if !errors.Is(err, ErrInvalidID) {
		t.Errorf("Esperava ErrInvalidID, mas obteve: %v", err)
	}
}

func TestUpdate_Success(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	// Criar uma personalidade
	createReq := &dto.CreatePersonalityRequest{
		Name:    "Alan Turing",
		History: "Matemático e cientista da computação britânico",
	}
	created, _ := service.Create(createReq)

	// Atualizar
	updateReq := &dto.UpdatePersonalityRequest{
		Name:    "Alan Mathison Turing",
		History: "Matemático, cientista da computação e criptoanalista britânico",
	}
	result, err := service.Update(created.ID, updateReq)

	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	if result.Name != updateReq.Name {
		t.Errorf("Esperava nome %s, mas obteve %s", updateReq.Name, result.Name)
	}
}

func TestDelete_Success(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	// Criar uma personalidade
	req := &dto.CreatePersonalityRequest{
		Name:    "Alan Turing",
		History: "Matemático e cientista da computação britânico",
	}
	created, _ := service.Create(req)

	// Deletar
	err := service.Delete(created.ID)

	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	// Verificar se foi deletado
	_, err = service.GetByID(created.ID)
	if !errors.Is(err, ErrPersonalityNotFound) {
		t.Error("Personalidade deveria ter sido deletada")
	}
}

func TestGetAll_Success(t *testing.T) {
	repo := newMockRepository()
	service := NewPersonalityService(repo)

	// Criar algumas personalidades
	personalities := []dto.CreatePersonalityRequest{
		{Name: "Alan Turing", History: "Matemático e cientista da computação"},
		{Name: "Ada Lovelace", History: "Matemática e escritora"},
		{Name: "Grace Hopper", History: "Cientista da computação e militar"},
	}

	for _, p := range personalities {
		service.Create(&p)
	}

	// Buscar todas
	result, err := service.GetAll()

	if err != nil {
		t.Errorf("Esperava sucesso, mas obteve erro: %v", err)
	}

	if len(result) != len(personalities) {
		t.Errorf("Esperava %d personalidades, mas obteve %d", len(personalities), len(result))
	}
}
