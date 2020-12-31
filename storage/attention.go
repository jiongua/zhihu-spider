package storage

import "github.com/satori/go.uuid"

type Attention struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	//复合索引
	UserID        uuid.UUID `gorm:"index:idx_member"`
	QuestionRefer int       `gorm:"index:idx_member"`
	Question      Question  `gorm:"foreignKey:QuestionRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
