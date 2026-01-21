package models

type Personality struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"unique;not null"`
	History string `json:"history" gorm:"type:text;not null"`
}

var Personalities []Personality
