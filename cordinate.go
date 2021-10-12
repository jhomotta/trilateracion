package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Cordinate struct {
	x    float64
	y    float64
	d    float64
	name string
}

var _ = godotenv.Load(".env")

func SateliteCordinate(name string, distance float64) Cordinate {
	var cor Cordinate
	switch name {
	case "kenobi":
		cor.x, err = strconv.ParseFloat(os.Getenv("kenobi_x"), 64)
		cor.y, err = strconv.ParseFloat(os.Getenv("kenobi_y"), 64)
	case "skywalker":
		cor.x, err = strconv.ParseFloat(os.Getenv("skywalker_x"), 64)
		cor.y, err = strconv.ParseFloat(os.Getenv("skywalker_y"), 64)
	case "sato":
		cor.x, err = strconv.ParseFloat(os.Getenv("sato_x"), 64)
		cor.y, err = strconv.ParseFloat(os.Getenv("sato_y"), 64)
	}

	cor.d = distance
	cor.name = name
	return cor

}
