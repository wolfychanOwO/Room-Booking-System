package booking

import (
	"context"
	"fmt"
	pb "hotelservice/proto/booking"
	"log"

	"github.com/jmoiron/sqlx"
	kafka "github.com/segmentio/kafka-go"
)

type Server struct {
	pb.UnimplementedBookingServiceServer
	DB *sqlx.DB
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (s *Server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
	log.Printf("Creating booking for hotel_id: %d and client_id: %d", req.HotelId, req.ClientId)

	writer := newKafkaWriter("kafka:9092", "bookings")
	defer writer.Close()
	msg := kafka.Message{
		Key:   []byte("test"),
		Value: []byte("test"),
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("produced")
	}

	return &pb.CreateBookingResponse{
		Message:   "Booking created successfully",
		BookingId: int32(1),
	}, nil
}

func (s *Server) GetBooking(ctx context.Context, req *pb.GetBookingRequest) (*pb.GetBookingResponse, error) {
	log.Printf("Getting booking information for booking_id: %d", req.BookingId)

	return &pb.GetBookingResponse{
		BookingId: req.BookingId,
		HotelId:   1,
		ClientId:  123,
		Status:    "confirmed",
	}, nil
}
