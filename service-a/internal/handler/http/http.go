package http

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	frequency   = 2 * time.Second
	frequencyMu sync.RWMutex
	validate    = validator.New()
)

func GetFrequency() time.Duration {
	frequencyMu.RLock()
	defer frequencyMu.RUnlock()
	return frequency
}

func SetFrequency(c echo.Context) error {
	type req struct {
		Seconds int `json:"seconds" validate:"required,gt=0,lte=3600"` // Max 1 jam
	}
	var body req
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Validasi payload
	if err := validate.Struct(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	frequencyMu.Lock()
	frequency = time.Duration(body.Seconds) * time.Second
	frequencyMu.Unlock()
	return c.JSON(http.StatusOK, map[string]string{"message": "frequency updated", "seconds": strconv.Itoa(body.Seconds)})
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

// SetupRoutes configures HTTP routes with rate limiting
func SetupRoutes(e *echo.Echo) {
	// Rate limiter: 10 requests/second per IP
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// Body limit: 1MB
	e.Use(middleware.BodyLimit("1M"))

	e.GET("/health", HealthCheck)
	e.GET("/frequency", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]int{"seconds": int(GetFrequency().Seconds())})
	})
	e.POST("/frequency", SetFrequency)
}
