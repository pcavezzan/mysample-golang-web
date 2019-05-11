package models

type Parking struct {
	Id 			int 	`json:"id,omitempty" db:"id"`
	Number      int 	`json:"number" db:"number"`
	Owner	string 	`json:"owner" db:"owner"`
	Car	string	`json:"car" db:"car"`
}
