package api

import "github.com/jinzhu/gorm"

// Handler ...
type Handler struct {
	Db *gorm.DB
}
