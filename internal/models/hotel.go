package models

type Hotel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Hotelier       int    `json:"hotelier_id"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `json:"address"`
}

