package models

type Booking struct { 
	ID int `json:"id"`
	ClientID int `json:"client_id"`
	HotelID int `json:"hotel_id"`
	RoomCategory string `json:"room_category"`
	Price int `json:"price"`
	Date string `json:"date"`

