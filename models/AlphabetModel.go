package models

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

func GetAlphabets() (*[]Alphabet, error) {

}
