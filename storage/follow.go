package storage

import "github.com/satori/go.uuid"

type Follow struct {
	FollowerID uuid.UUID `gorm:"primaryKey"`
	FolloweeID uuid.UUID `gorm:"primaryKey"`
}
