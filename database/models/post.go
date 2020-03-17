package models

import (
	"DataApi.Go/database/models/PV"
	"DataApi.Go/lib/common"
	"github.com/jinzhu/gorm"
)

// Post data model
type Post struct {
	gorm.Model
	Text       string        `sql:"type:text;"`
	StatPagePV PV.StatPagePV `gorm:"foreignkey:UserID"`
	UserID     uint
}

// Serialize serializes spark data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id":         p.ID,
		"text":       p.Text,
		"user":       p.StatPagePV.Serialize(),
		"created_at": p.CreatedAt,
	}
}
