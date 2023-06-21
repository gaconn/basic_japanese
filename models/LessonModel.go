package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	TypeLesson   int        `gorm:"not null"json:"type_lesson"` //1 is vocabulary - 2 is sentence case - 3 both
	Name         string     `gorm:"not null"json:"name"`
	Vietnamese   string     `json:"vietnamese"`
	WordList     []Word     `json:"word_list"`
	SentenceList []Sentence `json:"sentence_list"`
}

func (l *Lesson) GetLessons(page int) (*[]Lesson, error) {
	var list []Lesson
	offset := (page - 1) * NumberRecordLimit
	result := db.Limit(NumberRecordLimit).Offset(offset).Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}
	return &list, nil
}

func (l *Lesson) GetLesson(id int) error {
	result := db.Where("id=?", id).First(l)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Lesson) AddLesson() error {
	result := db.Create(l)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Lesson) DeleteLessons(list []*Lesson) error {
	result := db.Delete(list)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Lesson) UpdateLessons(id uint) error {
	result := db.Model(&Lesson{}).Where("id=?", id).Updates(l)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
