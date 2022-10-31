package v1

import (
	"net/http"

	getuser "base_service/internal/application/user/queries/get_user"
	httpmetrics "base_service/internal/metrics/http"
	"base_service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	v "github.com/gogovan-korea/ggx-kr-service-utils/validation"
)

type UserHandler struct {
	service *service.Service
	logger  logger.Logger
	metrics *httpmetrics.HttpMetrics
}

func NewUserHandler(service *service.Service, logger logger.Logger, metrics *httpmetrics.HttpMetrics) *UserHandler {
	return &UserHandler{service, logger, metrics}
}

// @BasePath /api/v1

// GetUser godoc
// @Tags Users
// @Summary Get user
// @Schemes
// @Description Get user by username
// @Accept json
// @Produce json
// @Param user body getuser.GetUserQuery true "User data"
// @Success 200 {object} getuser.User
// @Router /users [post]
func (h *UserHandler) GetUser(c *gin.Context) {
	var user getuser.GetUserQuery
	h.metrics.GetUserHttpRequests.Inc()

	if err := c.ShouldBind(&user); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, v.GetErrorMessage(fieldErr, h.logger))
		}

		c.JSON(http.StatusBadRequest, errors)
		return
	}

	userResponse, err := h.service.UserService.GetUserHandler.Handle(c.Request.Context(), &user)

	if err != nil {
		h.metrics.ErrorHttpRequests.Inc()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.metrics.SuccessHttpRequests.Inc()
	c.JSON(200, userResponse)
}

func (h *UserHandler) Test(c *gin.Context) {
	var result = 0
	for {
		result += 1
	}
	//c.JSON(200, result)
}
