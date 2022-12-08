package grpc

import (
	"context"

	pb "base_service/internal/api/grpc/proto_gen"
	createuser "base_service/internal/application/user/commands/create_user"
	getuser "base_service/internal/application/user/queries/get_user"

	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		s.metrics.GrpcRequestsLatency.WithLabelValues("GetUser").Observe(v)
	}))
	defer func() {
		timer.ObserveDuration()
	}()

	s.metrics.GetUserGrpcRequests.Inc()

	getUserQuery := &getuser.GetUserQuery{Username: in.GetUsername()}
	user, err := s.UserService.GetUserHandler.Handle(ctx, getUserQuery)

	if err != nil {
		s.metrics.ErrorGrpcRequests.Inc()
		return nil, err
	}

	s.metrics.SuccessGrpcRequests.Inc()

	return &pb.GetUserResponse{Username: user.Username, Email: user.Email, Phonenumber: user.PhoneNumber}, nil
}

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	createUserCommand := &createuser.CreateUserCommand{
		Username:    in.GetUsername(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhonenumber(),
	}
	success, err := s.UserService.CreateUserHandler.Handle(ctx, createUserCommand)

	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Sucucess: success}, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	//updateUserCommand := &updateuser.UpdateUserCommand{Username: in.GetUsername(), PhoneNumber: in.GetPhonenumber()}
	success, err := s.UserService.UpdateUserHandler.Handle(ctx)

	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{Sucucess: success}, nil
}
