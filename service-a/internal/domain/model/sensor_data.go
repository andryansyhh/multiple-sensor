package model

import "time"

type SensorData struct {
	SensorValue float64   `json:"sensor_value"` // Sensor reading value (float64)
	SensorType  string    `json:"sensor_type"`  // Type of sensor (fixed per instance)
	ID1         string    `json:"id1"`         	// Uppercase letter identifier
	ID2         int       `json:"id2"`         // Integer identifier
	Timestamp   time.Time `json:"timestamp"`   // ISO 8601 timestamp (UTC)
}