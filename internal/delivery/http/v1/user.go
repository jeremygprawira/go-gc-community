package v1

import (
	"go-gc-community/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userRoutes(api *gin.RouterGroup) {
	users := api.Group("/user")
	{
		users.POST("/register", h.userSignUp)
	}
}

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body userSignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) userSignUp(ctx *gin.Context) {
	var request registerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		//newResponse(c, http.StatusBadRequest, "invalid input body")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42202",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	user, err := h.usecase.Users.Register(usecase.RegisterUserRequest{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40001",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "User has been registered successfully",
		"userName": user.Name,
		"userEmail": user.Email,
		"userEncryptedPassword": user.Password,
	})
}