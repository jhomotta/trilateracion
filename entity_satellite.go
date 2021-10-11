package main

import "gorm.io/gorm"

type Satellite struct {
	gorm.Model
	Name     string `json:"name"`
	Distance int32  `json:"distance"`
	Message  string `json:"message"`
}
