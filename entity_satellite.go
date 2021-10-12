package main

import "gorm.io/gorm"

type Satellite struct {
	gorm.Model
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Message  string  `json:"message"`
}
