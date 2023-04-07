package controller

import "gorm.io/gorm"

type BaseController struct {
	db *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{db: db}
}


