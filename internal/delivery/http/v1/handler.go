package v1

import (
	"go-gc-community/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecases
}

func NewHandler (usecase *usecase.Usecases) *Handler{
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.userRoutes(v1)
	}
}