package models

import "time"

// Personality representa o modelo de personalidade no banco de dados
type Personality struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null;size:100"`
	History   string    `json:"history" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName especifica o nome da tabela no banco de dados
func (Personality) TableName() string {
	return "personalities"
}
