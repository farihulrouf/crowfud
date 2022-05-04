package handler
import (	
    "github.com/gin-gonic/gin"
    "net/http"
    "crowdgo/user"
)
type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	// map input dari user ke struct resiger input
	//strukct di atas kita passing sebagai parameter service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)
}