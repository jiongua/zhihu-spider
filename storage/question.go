package storage

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type (
	Topic struct {
		TopicID int `gorm:"primaryKey"`
		Name    string
	}

	Question struct {
		QuestionID    int `gorm:"primaryKey;autoIncrement"`
		OldQuestionID int
		TopicID       int
		Title         string
		AuthorID      uuid.UUID `gorm:"index"`
		Created       int       `gorm:"autoCreateTime"`
		Updated       int       `gorm:"autoUpdateTime"`
		Deleted       gorm.DeletedAt
	}
)

//func init() {
//	db := GetDB()
//	if err := db.AutoMigrate(&Topic{}, &Question{}); err != nil {
//		log.Println(err)
//	}
//}
