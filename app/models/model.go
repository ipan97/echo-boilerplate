package models

import (
"github.com/jinzhu/gorm"
"github.com/satori/go.uuid"
"time"
)
type Model struct {
	ID        string `gorm:"type:varchar(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) (er error) {
	uuidV4 := uuid.NewV4()
	return scope.SetColumn("ID", uuidV4.String())
}
