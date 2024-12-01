package main

import (
	"log"
	"net"

	"hotelservice/internal/hotel"
	pb "hotelservice/proto/hotel"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHotelServiceServer(s, &hotel.Server{})

	log.Println("Hotel Service started on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
