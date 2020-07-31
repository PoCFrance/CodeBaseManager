package models_v2

import "github.com/jinzhu/gorm"

type Module struct {
	gorm.Model
	Name      string     `gorm:"size:255;not null;unique" json:"name"`
	Path      string     `gorm:"size:255;not null;unique" json:"path"`
	Functions []Function `json:"functions"`
	Types     []Type     `json:"types"`
}
