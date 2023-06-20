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

var TypeWord = map[string]int{"HIRAGANA": 1, "HIRAGANA_COMBINE": 11, "KATAKANA": 2, "KATAKANA_COMBINE": 22, "KANJI": 3}

type ValidateError struct {
	Type    string
	Message string
	Code    int
}

func GetAllKatakana() (*[]Alphabet, error) {
	var list []Alphabet
	result := db.Where("type= ?", TypeWord["KATAKANA"]).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return &list, nil
}

func GetByIDKatakana(id int64) (*Alphabet, error) {
	var a Alphabet
	result := db.Where("type= ? AND id = ?", TypeWord["KATAKANA"], id).First(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	return &a, nil
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

func DeleteKatakana(id int64) (*Alphabet, error) {
	var alphabet = Alphabet{}
	result := db.Delete(&alphabet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &alphabet, nil
}

func UpdateKatakana(a *Alphabet) (*Alphabet, error) {
	result := db.Save(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
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
