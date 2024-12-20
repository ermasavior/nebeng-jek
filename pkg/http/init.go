package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nebeng-jek/pkg/configs"
	"nebeng-jek/pkg/http/middleware"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	pkgOtel "nebeng-jek/pkg/otel"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server struct {
	address string
	Router  *gin.Engine
}

func NewHTTPServer(cfg *configs.Config, otel *pkgOtel.OpenTelemetry) Server {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin.EnableJsonDecoderDisallowUnknownFields()

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.CorsHandler())
	router.Use(otelgin.Middleware(cfg.AppName))

	router.GET("/", healthCheck)
	router.GET(cfg.ApiPrefix+"/", healthCheck)
	router.GET(cfg.ApiPrefix+"/healthz", healthCheck)

	return Server{
		address: ":" + cfg.AppPort,
		Router:  router,
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse("Service is up and running."))
}

func (s *Server) Start(ctx context.Context) *http.Server {
	httpServer := &http.Server{
		Addr:         s.address,
		Handler:      s.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(ctx, "error running server", map[string]interface{}{logger.ErrorKey: err})
			return
		}
	}()
	logger.Info(ctx, fmt.Sprintf("server running on address: %s", s.address), nil)

	return httpServer
}
