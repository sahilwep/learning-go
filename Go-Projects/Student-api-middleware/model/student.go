/*
	// Struct Naming Convention:
		Uppercase: accessible outside this package.
		Lowercase: not Accessible outside this package

*/

package model

type Dob struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Addr struct {
	Pin     string `json:"pin,omitempty"` // might not always be provided. In that case, add "omitempty"
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address *Addr  `json:"address,omitempty"` // Using *Address with omitempty -> makes Location truly optional in JSON.
	Birth   Dob    `json:"birth"`
}
