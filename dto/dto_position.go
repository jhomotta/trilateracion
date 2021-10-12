package dto

type PositionDto struct {
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Message string `json:"message"`
}
