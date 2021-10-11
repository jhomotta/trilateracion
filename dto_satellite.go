package main

type SatelliteDto struct {
	Name     string   `json:"name"`
	Distance int32    `json:"distance"`
	Message  []string `json:"message"`
}
