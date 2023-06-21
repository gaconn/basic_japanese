package models

import "gorm.io/gorm"

type Sentence struct {
	gorm.Model
	Japanese   string  `gorm:"not null"json:"japanese"`
	Vietnamese string  `gorm:"not null"json:"vietnamese"`
	LessonID   uint    `json:"lesson_id"`
	Words      []*Word `gorm:"many2many:word_sentence;"json:"words"`
}

func (s *Sentence) GetSentences(page int) ([]*Sentence, error) {
	var list []*Sentence
	result := db.Offset(page * NumberRecordLimit).Limit(NumberRecordLimit).Find(list)

	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (s *Sentence) GetSentence(id int) error {
	result := db.Where("id=?", id).First(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Sentence) AddSentences(list []*Sentence) error {
	result := db.Create(list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Sentence) DeleteSentences(list []*Sentence) error {
	result := db.Delete(list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Sentence) UpdateSentences(id int) error {
	result := db.Model(&Sentence{}).Where("id=?", id).Updates(s)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
