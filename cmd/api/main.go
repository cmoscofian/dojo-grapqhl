package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}

func run(port string) error {
	// DefaultMeliRouter includes newrelic, datadog, attributes filter, jsonp and pprof middlewares.
	router := mlhandlers.DefaultMeliRouter()

	health := HealthChecker{}

	mapRoutes(router, health)

	return router.Run(":" + port)
}

func mapRoutes(r *gin.Engine, health HealthChecker) {
	r.GET("/ping", health.PingHandler)
}

// HealthChecker struct provides the handler for a health check endpoint.
// This is a starting point which allows each application to
// customize their health check strategy.
type HealthChecker struct{}

// PingHandler returns a successful pong answer to all HTTP requests.
func (h HealthChecker) PingHandler(c *gin.Context) {
	if txn := nrgin.Transaction(c); txn != nil {
		txn.Ignore()
	}

	c.String(http.StatusOK, "pong")
}
