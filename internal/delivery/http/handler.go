package http

import (
	"fmt"
	"go-gc-community/docs"
	"go-gc-community/internal/config"
	v1 "go-gc-community/internal/delivery/http/v1"
	"go-gc-community/internal/usecase"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.Http.Host
	}

	if cfg.Environment != config.Prod {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Init router
	router.GET("/health-check", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.usecase)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
