package getuser

import (
	interfaces "base_service/internal/domain/interfaces/user"
	"context"
	"github.com/gogovan-korea/ggx-kr-service-utils/tracing"
	// Just exmple, in the real world, using https://github.com/gogovan-korea/s14e-backend-proto
	proto "base_service/internal/api/grpc/proto_gen"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
)

type GetUserHandler struct {
	logger            logger.Logger
	repo              interfaces.UserQueryRepository
	redisRepo         interfaces.CacheRepository
	userServiceClient proto.UserServiceClient
}

func NewGetUserHandler(
	logger logger.Logger,
	repo interfaces.UserQueryRepository,
	redisRepo interfaces.CacheRepository,
	userServiceClient proto.UserServiceClient,
) *GetUserHandler {
	return &GetUserHandler{logger, repo, redisRepo, userServiceClient}
}

func (h *GetUserHandler) Handle(ctx context.Context, getUserQuery *GetUserQuery) (user *User, err error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "GetUserHandler.Handle")
	defer func() {
		span.RecordError(err)
		span.End()
	}()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx)
	resp, err := h.userServiceClient.CreateUser(ctx, &proto.CreateUserRequest{Username: getUserQuery.Username})
	if err != nil {
		return nil, err
	}
	print(resp)
	// Get from cache first
	if user := h.redisRepo.GetUser(ctx, getUserQuery.Username); user != nil {
		u := User(*user)
		return &u, nil
	}

	userData, err := h.repo.GetUser(ctx, getUserQuery.Username)

	if err != nil {
		return nil, err
	}
	return &User{
		Username:    userData.Username,
		Email:       userData.Email,
		PhoneNumber: userData.PhoneNumber,
	}, nil
}
