package models

type Parking struct {
	Number      int 	`json:"number"`
	Owner	string 	`json:"owner"`
	Car	string	`json:"car"`
}
