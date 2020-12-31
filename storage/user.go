package storage

import (
	"github.com/satori/go.uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	HeadLine      string
	Name          string
	AnswerCount   int
	FolloweeCount int
	FollowerCount int
	VoteCount     int
}
