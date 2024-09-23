package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"nebeng-jek/pkg/http/middleware"
	otelMiddleware "nebeng-jek/pkg/http/middleware/otelmetrics"
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

func NewHTTPServer(appName, appEnv, appPort string, otel *pkgOtel.OpenTelemetry) Server {
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gin.EnableJsonDecoderDisallowUnknownFields()

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.CorsHandler())
	router.Use(otelgin.Middleware(appName))
	router.Use(middleware.TracerIDHandler())
	router.Use(otelMiddleware.Middleware(appName, otel.Meter))

	router.GET("/", healthCheck)
	router.GET("/healthz", healthCheck)

	return Server{
		address: ":" + appPort,
		Router:  router,
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse("Service is up and running."))
}

func (s *Server) Start() *http.Server {
	httpServer := &http.Server{
		Addr:         s.address,
		Handler:      s.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(context.Background(), "error running server", map[string]interface{}{logger.ErrorKey: err})
		}
	}()
	logger.Info(context.Background(), fmt.Sprintf("server running on address: %s", s.address), nil)

	return httpServer
}
