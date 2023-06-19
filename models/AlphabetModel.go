package models

import (
	"github.com/quan12xz/basic_japanese/utils"
)

type Alphabet struct {
	ID           int        `json:"id"`
	Word         string     `json:"word"`
	Stroke       string     `json:"stroke"`
	Radical      string     `json:"radical"`
	Meaning      string     `json:"meaning"`
	Detail       string     `json:"detail"`
	Note         string     `json:"note"`
	IsImportance string     `json:"is_importance"`
	ImageUrl     string     `json:"image_url"`
	SimilarWord  []Alphabet `json:"similar_word"`
	SimilarSound []Alphabet `json:"similar_sound"`
	VariantList  []Alphabet `json:"variant"`
}

var db = utils.DBUtilInstance.DB

func GetAllKatakana() (*[]Alphabet, error) {
	var list []Alphabet
	db.Find(&list)

	return &list, nil
}

func CreateKatakana(alphabet *Alphabet) (*Alphabet, error) {
	result := db.Create(alphabet)
	if result.Error != nil {
		return nil, result.Error
	}
	return alphabet, nil
}
