package usecase

import (
	"math/rand"
	"service-a/internal/domain/model"
	"time"
)

// DataGenerator encapsulates logic for generating sensor data.
type DataGenerator struct {
	sensorType string
}

// NewDataGenerator creates a new DataGenerator with a fixed sensor type.
func NewDataGenerator(sensorType string) *DataGenerator {
	return &DataGenerator{sensorType: sensorType}
}

// Generate creates a new SensorData instance with random values.
func (g *DataGenerator) Generate() *model.SensorData {
	rand.Seed(time.Now().UnixNano())

	return &model.SensorData{
		SensorValue: rand.Float64() * 100,              // Random float between 0 and 100
		SensorType:  g.sensorType,                      // Fixed sensor type
		ID1:         string(rune('A' + rand.Intn(26))), // Random uppercase letter A-Z
		ID2:         rand.Intn(100) + 1,                // Random integer 1-100
		Timestamp:   time.Now().UTC(),                  // Current time in UTC
	}
}
