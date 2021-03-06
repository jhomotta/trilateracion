package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"meradolibre.com/q/v2/trilateration"

	"meradolibre.com/q/v2/dto"
)

type (
	SatelliteApi struct{}
)

func NewSatelliteApi() *SatelliteApi {
	return &SatelliteApi{}
}

func CreateSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items dto.SatelliteDto
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		panic(err)
	}

	var cons []trilateration.Cordinate

	for index, element := range items.Satellite {

		log.Println(index)

		var con trilateration.Cordinate
		con = trilateration.SateliteCordinate(element.Name, element.Distance)
		cons = append(cons, con)
	}

	x, y := cons[0].Trilaterate(cons[1], cons[2])
	log.Println("Cordinates of intersections center using formula x:", x, "y:", y)

	var secretMessage string

	for i, a := range items.Satellite[0].Message {
		if a != "" {
			secretMessage = secretMessage + " " + a
		} else if items.Satellite[1].Message[i] != "" {
			secretMessage = secretMessage + " " + items.Satellite[1].Message[i]
		} else {
			secretMessage = secretMessage + " " + items.Satellite[2].Message[i]
		}
	}

	var position dto.PositionDto
	position.Position.X = x
	position.Position.Y = y
	position.Message = secretMessage

	log.Println(cons)
	json.NewEncoder(w).Encode(position)
}
