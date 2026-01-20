package models

type Personality struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"unique;not null"`
	History string `json:"history" gorm:"type:text"`
}

var Personalities []Personality
