package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDataGenerator(t *testing.T) {
	gen := NewDataGenerator("temperature")
	assert.Equal(t, "temperature", gen.sensorType, "sensorType should be temperature")
}

func TestGenerate(t *testing.T) {
	gen := NewDataGenerator("humidity")
	data := gen.Generate()

	assert.True(t, data.SensorValue >= 0 && data.SensorValue <= 100, "SensorValue should be between 0 and 100")
	assert.Equal(t, "humidity", data.SensorType, "SensorType should be humidity")
	assert.True(t, data.ID1 >= "A" && data.ID1 <= "Z", "ID1 should be uppercase letter")
	assert.True(t, data.ID2 >= 1 && data.ID2 <= 100, "ID2 should be between 1 and 100")
	assert.False(t, data.Timestamp.After(time.Now().UTC()), "Timestamp should not be in future")
}
