package repository

import (
	"go-api-rest/models"

	"gorm.io/gorm"
)

// PersonalityRepository define a interface para operações de dados
type PersonalityRepository interface {
	Create(personality *models.Personality) error
	FindAll() ([]models.Personality, error)
	FindByID(id uint) (*models.Personality, error)
	Update(personality *models.Personality) error
	Delete(id uint) error
	ExistsByName(name string) (bool, error)
}

// personalityRepository implementa PersonalityRepository
type personalityRepository struct {
	db *gorm.DB
}

// NewPersonalityRepository cria uma nova instância do repositório
func NewPersonalityRepository(db *gorm.DB) PersonalityRepository {
	return &personalityRepository{db: db}
}

func (r *personalityRepository) Create(personality *models.Personality) error {
	return r.db.Create(personality).Error
}

func (r *personalityRepository) FindAll() ([]models.Personality, error) {
	var personalities []models.Personality
	err := r.db.Order("id ASC").Find(&personalities).Error
	return personalities, err
}

func (r *personalityRepository) FindByID(id uint) (*models.Personality, error) {
	var personality models.Personality
	err := r.db.First(&personality, id).Error
	if err != nil {
		return nil, err
	}
	return &personality, nil
}

func (r *personalityRepository) Update(personality *models.Personality) error {
	return r.db.Save(personality).Error
}

func (r *personalityRepository) Delete(id uint) error {
	return r.db.Delete(&models.Personality{}, id).Error
}

func (r *personalityRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Personality{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
