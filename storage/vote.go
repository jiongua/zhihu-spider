package storage

import "github.com/satori/go.uuid"

type VoteAnswer struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	UserID      uuid.UUID
	AnswerRefer int
	Answer      Answer `gorm:"foreignKey:AnswerRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type VoteComment struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement"`
	UserID       uuid.UUID
	CommentRefer uint64
	Comment      Comment `gorm:"foreignKey:CommentRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
