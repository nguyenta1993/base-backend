package getuser

import (
	"base_service/internal/domain/entities"
	mock_user "base_service/internal/domain/interfaces/user/mocks"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	rootCtx := context.Background()
	span, ctx := opentracing.StartSpanFromContext(rootCtx, "TestHandle")
	defer span.Finish()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserQueryRepository := mock_user.NewMockUserQueryRepository(mockCtrl)
	mockUserQueryRepository.EXPECT().GetUser(ctx, "k").Return(&entities.User{
		Username:    "k",
		Email:       "k@gmail.com",
		PhoneNumber: "123456789",
	}, nil)

	mockCacheRepository := mock_user.NewMockCacheRepository(mockCtrl)
	// mockCacheRepository.EXPECT().GetUser(ctx, "k").Return(&entities.User{
	// 	Username:    "k",
	// 	Email:       "k@gmail.com",
	// 	PhoneNumber: "123456789",
	// })

	mockCacheRepository.EXPECT().GetUser(ctx, "k").Return(nil)

	handler := GetUserHandler{
		repo:      mockUserQueryRepository,
		redisRepo: mockCacheRepository,
	}

	user, _ := handler.Handle(rootCtx, &GetUserQuery{
		Username: "k",
	})

	assertions := require.New(t)
	assertions.Equal("k", user.Username, "wrong username")
}
