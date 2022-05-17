package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title" example:"Book A"`
	Author    string `json:"author" example:"Dino"`
	Publisher string `json:"publisher" example:"Creative Company"`
	Name      string `json:"name" example:"Creative Company"`
}
