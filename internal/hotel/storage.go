package hotel

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

type Hotel struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Hotelier int    `json:"hotelier"`
	Rating   int    `json:"rating"`
	Country  string `json:"country"`
	Address  string `json:"address"`
}

func NewStorage(conn string) *Storage {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic("Connection failed: " + err.Error())
	}
	return &Storage{db: db}
}

func (s *Storage) GetHotels() ([]Hotel, error) {
	rows, err := s.db.Query("SELECT ID, Name, Hotelier, Rating, Country, Address FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []Hotel
	for rows.Next() {
		var hotel Hotel
		if err := rows.Scan(&hotel.ID, &hotel.Name, &hotel.City); err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (s *Storage) AddHotel(hotel Hotel) error {
	_, err := s.db.Exec(
		`INSERT INTO hotels (
			ID, 
			Name,
			Hotelier, 
			Rating, 
			Country, 
			Address
		) VALUES ($1, $2, $3, $4, $5, $6)`,
		hotel.ID,
		hotel.Name,
		hotel.Hotelier,
		hotel.Rating,
		hotel.Country,
		hotel.Address,
	)
	return err
}
