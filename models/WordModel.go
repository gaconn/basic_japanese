package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	TypeWord     int        `gorm:"not null"json:"type_word"` // 1-hiragana 2-katakana 3-kanji
	Japanese     string     `gorm:"not null"json:"japanese"`
	Vietnamese   string     `gorm:"not null"json:"vietnamese"`
	Category     string     `json:"category"` // adj, adv, n, v...
	LessonID     uint       `json:"lesson_id"`
	SentenceList []Sentence `gorm:"many2many:word_sentence;"json:"sentence_list"`
}

func (w *Word) GetWords(page int) ([]*Word, error) {
	var list []*Word
	result := db.Offset(page * NumberRecordLimit).Limit(NumberRecordLimit).Find(list)

	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (w *Word) GetWord(id int) error {
	result := db.Where("id=?", id).First(w)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *Word) AddWords(list []*Word) error {
	result := db.Create(list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *Word) DeleteWords(list []*Word) error {
	result := db.Delete(list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (w *Word) UpdateWord(id int) error {
	result := db.Model(&Word{}).Where("id=?", id).Updates(w)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
