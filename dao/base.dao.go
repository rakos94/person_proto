package dao

import (
	"gorm.io/gorm"
)

var g *gorm.DB

// SetDao ...
func SetDao(dg *gorm.DB) {
	g = dg
}
