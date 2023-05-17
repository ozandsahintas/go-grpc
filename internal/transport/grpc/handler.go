package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"odesch.com/odesch/grpc-app/internal/rocket"
	r "odesch.com/odesch/grpc-app/protos/rocket/v1"
)

type RocketService interface {
	GetRocketById(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, r rocket.Rocket) (rocket.Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
}

// Handler - will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("could not listen on port 50051: %s", err)
		return err
	}

	grpcServer := grpc.NewServer()
	r.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *r.GetRocketRequest) (*r.GetRocketResponse, error) {
	log.Println("Get Rocket gRPC endpoint hit")

	rocket, err := h.RocketService.GetRocketById(ctx, req.Id)
	if err != nil {
		log.Println("failed to retrieve rocket by this id")
		return &r.GetRocketResponse{}, err
	}

	return &r.GetRocketResponse{
		R: &r.Rocket{
			Id:   rocket.Id,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *r.AddRocketRequest) (*r.AddRocketResponse, error) {
	log.Println("Add Rocket gRPC endpoint hit")

	rocket, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		Id:   req.R.Id,
		Name: req.R.Name,
		Type: req.R.Type,
	})
	if err != nil {
		log.Println("failed to add rocket")
		return &r.AddRocketResponse{}, err
	}

	return &r.AddRocketResponse{
		R: &r.Rocket{
			Id:   rocket.Id,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

func (h Handler) DeleteRocket(ctx context.Context, req *r.DeleteRocketRequest) (*r.DeleteRocketResponse, error) {
	log.Println("Delete Rocket gRPC endpoint hit")
	err := h.RocketService.DeleteRocket(ctx, req.R.Id)
	if err != nil {
		return &r.DeleteRocketResponse{
			Status: "failed",
		}, err
	}
	return &r.DeleteRocketResponse{
		Status: "success",
	}, nil
}
