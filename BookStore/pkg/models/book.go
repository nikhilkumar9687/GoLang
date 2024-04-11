package models

import (
	"github.com/jinzhu/gorm"
	"./pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm :""json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
}

func init () {
	config.Connect()
}