package test

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ServerImpl is awesome
type serverImpl struct{}

// Register is awesome
func Register(g *grpc.Server) {
	RegisterHelloServer(g, &serverImpl{})
}

// SayHi is awesome
func (s *serverImpl) SayHi(ctx context.Context, req *SayHiRequest) (*SayHiResponse, error) {

	if req.Name == "" {
		req.Name = "World"
	} else if strings.Contains(req.Name, " ") {
		return nil, status.Errorf(codes.InvalidArgument, "`Name` must be a single word ... No spaces")
	}

	return &SayHiResponse{Message: fmt.Sprintf("Hello %s!", req.Name)}, nil
}
