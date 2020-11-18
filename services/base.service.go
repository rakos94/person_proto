package services

import (
	"gorm.io/gorm"
)

var g *gorm.DB

// SetService ...
func SetService(dg *gorm.DB) {
	g = dg
}
