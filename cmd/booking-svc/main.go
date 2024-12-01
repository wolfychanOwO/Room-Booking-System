package main

import (
	"log"
	"net"
	"net/http"

	"hotelservice/internal/booking"
	pb "hotelservice/proto/booking"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	bookingService := &booking.Server{}

	pb.RegisterBookingServiceServer(grpcServer, bookingService)

	r := mux.NewRouter()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	bookingClient := pb.NewBookingServiceClient(conn)
	handler := booking.NewHandler(bookingClient)

	r.HandleFunc("/bookings", handler.GetBookings).Methods("GET")
	r.HandleFunc("/bookings", handler.AddBooking).Methods("POST")

	go func() {
		log.Println("HTTP server started on port 8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatalf("failed to serve HTTP: %v", err)
		}
	}()

	log.Println("Booking Service started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
