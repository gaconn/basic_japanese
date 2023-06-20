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
type ValidateError struct {
	Type    string
	Message string
	Code    int
}

func GetAllKatakana() (*[]Alphabet, error) {
	var list []Alphabet
	db.Find(&list)

	return &list, nil
}

func CreateKatakana(alphabet *Alphabet) (*Alphabet, error) {
	if db == nil {
		return nil, nil
	}
	result := db.Create(alphabet)
	if result.Error != nil {
		return nil, result.Error
	}
	return alphabet, nil
}

func ValidateAlphabet(ap *Alphabet) *[]ValidateError {
	var err []ValidateError
	if ap.Word == "" {
		err = append(err, ValidateError{Type: "Blank", Message: "Word can't be empty", Code: 1})
	}

	if ap.Stroke == 0 {
		err = append(err, ValidateError{Type: "Blank", Message: "Stroke can't be empty", Code: 1})
	}

	if ap.Meaning == "" {
		err = append(err, ValidateError{Type: "Blank", Message: "Meaning can't be empty", Code: 1})
	}

	return &err
}
