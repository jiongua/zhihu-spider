package storage

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Comment struct {
	CommentID    uint64 `gorm:"primaryKey"`
	OldCommentID uint64
	AnswerRefer  int    `gorm:"index"`
	Answer       Answer `gorm:"foreignKey:AnswerRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content      string
	AuthorID     uuid.UUID `gorm:"index"`
	VoteCount    int
	Created      int `gorm:"autoCreateTime"`
	Deleted      gorm.DeletedAt
}

//type ChildComment struct {
//	CommentID uint64 `gorm:"primaryKey"`
//	ParentID uint64
//	Content string
//	AuthorID uuid.UUID `gorm:"index"`
//	ParentAuthorID uuid.UUID
//	VoteCount int
//	Created int	`gorm:"autoCreateTime"`
//	Deleted gorm.DeletedAt
//}
