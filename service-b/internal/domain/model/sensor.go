package model

import "time"

type SensorData struct {
	ID          int64     `db:"id"`
	SensorValue float64   `db:"sensor_value"`
	SensorType  string    `db:"sensor_type"`
	ID1         string    `db:"id1"`
	ID2         int       `db:"id2"`
	Timestamp   time.Time `db:"timestamp"`
}
