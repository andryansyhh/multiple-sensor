package http

import (
	"net/http"
	"strconv"
	"time"

	"service-b/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var validate = validator.New()

type SensorController struct {
	usecase usecase.SensorUsecase
}

func NewSensorController(usecase usecase.SensorUsecase) *SensorController {
	return &SensorController{usecase: usecase}
}

func (ctrl *SensorController) GetData(c echo.Context) error {
	// ambil query params
	id1 := c.QueryParam("id1")
	id2Str := c.QueryParam("id2")
	startStr := c.QueryParam("start")
	endStr := c.QueryParam("end")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	var (
		startTime *time.Time
		endTime   *time.Time
		id2Int    int
		limitInt  int
		offsetInt int
	)

	// parse optional id2
	if id2Str != "" {
		val, err := strconv.Atoi(id2Str)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id2"})
		}
		id2Int = val
	}

	// parse optional start
	if startStr != "" {
		t, err := time.Parse(time.RFC3339, startStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid start time"})
		}
		startTime = &t
	}

	// parse optional end
	if endStr != "" {
		t, err := time.Parse(time.RFC3339, endStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid end time"})
		}
		endTime = &t
	}

	// parse limit
	if limitStr != "" {
		val, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid limit"})
		}
		limitInt = val
	}

	// parse offset
	if offsetStr != "" {
		val, err := strconv.Atoi(offsetStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid offset"})
		}
		offsetInt = val
	}

	// panggil usecase
	data, err := ctrl.usecase.Query(id1, id2Int, startTime, endTime, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

// Update sensor by ID
func (ctrl *SensorController) UpdateSensor(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var body struct {
		SensorValue float64 `json:"sensor_value"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	if err := ctrl.usecase.UpdateByID(id, body.SensorValue); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "updated successfully"})
}

// Delete sensor by ID
func (ctrl *SensorController) DeleteSensor(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	if err := ctrl.usecase.DeleteByID(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "deleted successfully"})
}

func (ctrl *SensorController) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}

func SetupRoutes(e *echo.Echo, ctrl *SensorController) {
	// Rate limiter: 10 requests/second per IP
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	// Body limit: 1MB (meski GET, untuk safety)
	e.Use(middleware.BodyLimit("1M"))

	e.GET("/health", ctrl.HealthCheck)
	e.GET("/data", ctrl.GetData)
	e.PUT("/data/:id", ctrl.UpdateSensor)
	e.DELETE("/data/:id", ctrl.DeleteSensor)

}
