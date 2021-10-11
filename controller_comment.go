package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strings"
)

type (
	// CartController represents the controller for create cart and calculate tax
	SatelliteController struct{}
)

func NewSatelliteController() *SatelliteController {
	return &SatelliteController{}
}

func GetSatellites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var satellites []Satellite
	DB.Find(&satellites)

	satellitesdto := []SatelliteDto{}

	for index, element := range satellites {

		log.Println(index)
		satellitedto := SatelliteDto{
			Name:     element.Name,
			Distance: element.Distance,
			Message:  strings.Split(element.Message, ","),
		}
		satellitesdto = append(satellitesdto, satellitedto)

	}

	/*	coral := []string{"blue coral", "staghorn coral", "pillar coral", "", "elkhorn coral"}
		satellitesdto1 := SatelliteDto{
			Name:     "aaaaa",
			Distance: 11,
			Message:  coral,
		}
		satellitesdto2 := SatelliteDto{
			Name:     "bbbbb",
			Distance: 22,
			Message:  coral,
		}
		satellitesdto3 := SatelliteDto{
			Name:     "ccccc",
			Distance: 33,
			Message:  coral,
		}*/
	calc()
	json.NewEncoder(w).Encode(satellites)
}

func CreateSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item SatelliteDto
	json.NewDecoder(r.Body).Decode(&item)

	var satellite Satellite

	satellite.Name = item.Name
	satellite.Message = strings.Join(item.Message, ",")
	satellite.Distance = item.Distance
	log.Println(satellite)

	DB.Create(&satellite)
	json.NewEncoder(w).Encode(satellite)
}

func calc() {
	//punto a
	var xa float64 = 3.0
	var ya float64 = 5.0
	var ra float64 = 7.62
	//punto b
	var xb float64 = -2.0
	var yb float64 = -3.0
	var rb float64 = 5.385
	//punto c
	//var xc float64 = 4.0
	//var yc float64 = -4.0
	//var rc float64 = 10

	var dac float64 = math.Sqrt(math.Pow((xa-xb), 2) + math.Pow((ya-yb), 2))

	var x = (math.Pow(ra, 2) - math.Pow(rb, 2) + math.Pow(dac, 2)) / 2 * dac
	//punto d
	//var xd float64 = 0.0
	//var yd float64 = 0.0

	//ca, cb y cc = xd^2 + yd^2

	//var xd float64 = ((-(math.Pow(xb, 2))-(math.Pow(yb, 2))+(math.Pow(xc, 2))+(math.Pow(yc, 2))-db+dc)*(ya-yb) - (-(math.Pow(xb, 2))-(math.Pow(yb, 2))+(math.Pow(xa, 2))+(math.Pow(ya, 2))-db+da)*(yc-yb)) / ((xb-xa)*(yc-yb) - (xa-xc)*(ya-yb))
	//var xda = ((xb-xa)*(yc-yb) - (xa-xc)*(ya-yb)) / ((ya - yb) * (yc - yb))
	//var xdb = ((-(math.Pow(xb, 2))-(math.Pow(yb, 2))+(math.Pow(xc, 2))+(math.Pow(yc, 2))-math.Pow(db, 2)+math.Pow(dc, 2))*(ya-yb) - (-(math.Pow(xb, 2))-(math.Pow(yb, 2))+(math.Pow(xa, 2))+(math.Pow(ya, 2))-math.Pow(db, 2)+math.Pow(da, 2))*(yc-yb)) / ((yc - yb) * (ya - yb))
	//log.Println(xda)
	//log.Println(xdb)
	log.Println(dac)
	log.Println(x)
}

/*
func GetSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var catellite Catellite
	DB.First(&catellite, params["id"])
	json.NewEncoder(w).Encode(catellite)
}



func UpdateSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var catellite Catellite
	DB.First(&catellite, params["id"])
	json.NewDecoder(r.Body).Decode(&catellite)
	DB.Save(&catellite)
	json.NewEncoder(w).Encode(catellite)
}

func DeleteSatellite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Catellite Catellite
	DB.Delete(&Catellite, params["id"])
	json.NewEncoder(w).Encode(Catellite)
}
*/
