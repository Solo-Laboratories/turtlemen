package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/Solo-Laboratories/turtlemen/auth-mock/mockauth" // Import the generated files from the same directory

	"github.com/urfave/cli/v3"
	"google.golang.org/grpc"
)

func main() {
	cmd := &cli.Command{
		Name:  "auth-server",
		Usage: "Authentication Server for Turtlemen Application.",
		Action: func(context.Context, *cli.Command) error {
			start()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Info(err.Error())
	}
}

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) GetRandomString(ctx context.Context, req *pb.AuthRequest) (*pb.RandomStringResponse, error) {
	// Generate a random string
	rand.Seed(time.Now().UnixNano())
	randomString := fmt.Sprintf("random-%d", rand.Intn(100000))

	// Convert to JSON
	jsonResponse, err := json.Marshal(map[string]string{"random_string": randomString})
	if err != nil {
		return nil, err
	}

	return &pb.RandomStringResponse{RandomString: string(jsonResponse)}, nil
}

func start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error(fmt.Sprintf("failed to listen: %v", err))
		panic("failed to start gRPC server")
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})

	slog.Info("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		slog.Error(fmt.Sprintf("failed to serve: %v", err))
		panic("failed to serve gRPC server")
	}
}
