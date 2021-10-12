package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	SatelliteController struct{}
)

func NewSatelliteController() *SatelliteController {
	return &SatelliteController{}
}

func CreateSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items SatelliteDto
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		panic(err)
	}

	var cons []Cordinate

	for index, element := range items.Satellite {

		log.Println(index)
		var con Cordinate
		con = SateliteCordinate(element.Name, element.Distance)
		cons = append(cons, con)
	}

	x, y := cons[0].Trilaterate(cons[1], cons[2])
	log.Println("Cordinates of intersections center using formula x:", x, "y:", y)

	var position PositionDto
	position.Position.X = x
	position.Position.Y = y
	log.Println(cons)
	json.NewEncoder(w).Encode(position)
}
