package main

type SatelliteDto struct {
	Satellite []struct {
		Name     string   `json:"name"`
		Distance float64  `json:"distance"`
		Message  []string `json:"message"`
	} `json:"satellite"`
}
