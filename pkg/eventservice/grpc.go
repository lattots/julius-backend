package eventservice

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/lattots/julius/pkg/event"
	pb "github.com/lattots/julius/proto"
)

type Server struct {
	pb.UnimplementedEventServiceServer
	db         *sql.DB
	port       int
	gRPCServer *grpc.Server
}

func New(db *sql.DB, port int) *Server {
	return &Server{db: db, port: port}
}

func (s *Server) ListenAndServe() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}

	s.gRPCServer = grpc.NewServer()

	pb.RegisterEventServiceServer(s.gRPCServer, s)

	if err = s.gRPCServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetEvent(ctx context.Context, r *pb.SingleEventRequest) (*pb.SingleEventResponse, error) {
	e := event.New(s.db)
	err := e.GetByID(int(r.EventID))
	if err != nil {
		return nil, err
	}

	price := float32(e.Price)

	res := pb.SingleEventResponse{
		EventID:    uint32(e.ID),
		Name:       e.Name,
		Host:       e.Host,
		StartTime:  timestamppb.New(e.Start),
		EndTime:    timestamppb.New(e.End),
		DressCode:  &e.DressCode,
		Theme:      &e.Theme,
		Price:      &price,
		SignupLink: &e.SignupLink,
	}

	return &res, nil
}

func (s *Server) GetEvents(ctx context.Context, r *pb.MultiEventRequest) (*pb.MultiEventResponse, error) {
	log.Printf("GetEvents called with count:%d\n", r.Count)
	return nil, nil
}
