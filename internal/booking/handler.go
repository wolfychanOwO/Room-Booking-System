package booking

import (
	"context"
	"encoding/json"
	pb "hotelservice/proto/booking"
	"net/http"
)

type Handler struct {
	client pb.BookingServiceClient
}

func NewHandler(client pb.BookingServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) GetBookings(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	bookings, err := h.client.GetBooking(ctx, &pb.GetBookingRequest{BookingId: 1})
	if err != nil {
		http.Error(w, "Error fitching bookings: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func (h *Handler) AddBooking(w http.ResponseWriter, r *http.Request) {
	var booking Booking
	ctx := context.Background()
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid JSON input: "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err := h.client.CreateBooking(ctx, &pb.CreateBookingRequest{HotelId: 1, ClientId: 1})
	if err != nil {
		http.Error(w, "Error adding booking: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Booking added"))
}
