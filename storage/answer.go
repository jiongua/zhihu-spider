package storage

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Answer struct {
	AnswerID      int `gorm:"primaryKey"`
	OldAnswerID   int
	QuestionRefer int      `gorm:"index"`
	Question      Question `gorm:"foreignKey:QuestionRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Excerpt       string
	Content       string
	AuthorID      uuid.UUID `gorm:"index"`
	CommentCount  int
	VoteCount     int
	Created       int `gorm:"autoCreateTime"`
	Updated       int `gorm:"autoUpdateTime"`
	Deleted       gorm.DeletedAt
}
