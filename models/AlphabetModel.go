package models

import (
	"gorm.io/gorm"
)

type Alphabet struct {
	gorm.Model
	Type         int        `gorm:"not null"json:"type"`
	Word         string     `gorm:"not null"json:"word"`
	Stroke       int        `json:"stroke"`
	Meaning      string     `json:"meaning"`
	Detail       string     `json:"detail"`
	Note         string     `json:"note"`
	IsImportance string     `json:"is_importance"`
	ImageUrl     string     `json:"image_url"`
	SimilarWord  []Alphabet `gorm:"many2many:similar_word"json:"similar_word"`
	SimilarSound []Alphabet `gorm:"many2many:similar_sound"json:"similar_sound"`
	VariantList  []Alphabet `gorm:"many2many:variant"json:"variant"`
}

func UpdateWord(a *Alphabet) (*Alphabet, error) {
	result := db.Save(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

func GetWordByType(t int) (*[]Alphabet, error) {
	var list []Alphabet
	result := db.Where("type= ?", t).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return &list, nil
}

func GetWordByID(id int64) (*Alphabet, error) {
	var a Alphabet
	result := db.Where("id = ?", id).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return &a, nil
}
