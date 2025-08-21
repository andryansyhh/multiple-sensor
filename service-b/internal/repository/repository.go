package repository

import (
	"database/sql"
	"service-b/internal/domain/model"
	"time"
)

type SensorRepository interface {
	Save(data *model.SensorData) error
	Query(id1 string, id2 int, start, end *time.Time, limit, offset int) ([]model.SensorData, error)
	UpdateByID(id int64, newValue float64) error
	DeleteByID(id int64) error
}

type sensorRepository struct {
	db *sql.DB
}

func NewSensorRepository(db *sql.DB) SensorRepository {
	return &sensorRepository{db: db}
}

func (r *sensorRepository) Save(data *model.SensorData) error {
	_, err := r.db.Exec(`
		INSERT INTO sensors(sensor_value, sensor_type, id1, id2, timestamp)
		VALUES (?, ?, ?, ?, ?)`,
		data.SensorValue, data.SensorType, data.ID1, data.ID2, data.Timestamp)
	return err
}

func (r *sensorRepository) Query(id1 string, id2 int, start, end *time.Time, limit, offset int) ([]model.SensorData, error) {
	query := `SELECT id, sensor_value, sensor_type, id1, id2, timestamp FROM sensors WHERE 1=1`
	args := []interface{}{}

	if id1 != "" {
		query += " AND id1 = ?"
		args = append(args, id1)
	}
	if id2 > 0 {
		query += " AND id2 = ?"
		args = append(args, id2)
	}
	if start != nil && end != nil {
		query += " AND timestamp BETWEEN ? AND ?"
		args = append(args, *start, *end)
	}
	query += " ORDER BY timestamp DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.SensorData
	for rows.Next() {
		var d model.SensorData
		if err := rows.Scan(&d.ID, &d.SensorValue, &d.SensorType, &d.ID1, &d.ID2, &d.Timestamp); err != nil {
			return nil, err
		}
		results = append(results, d)
	}
	return results, nil
}

func (r *sensorRepository) UpdateByID(id int64, newValue float64) error {
	_, err := r.db.Exec(`UPDATE sensors SET sensor_value=? WHERE id=?`, newValue, id)
	return err
}

func (r *sensorRepository) DeleteByID(id int64) error {
	_, err := r.db.Exec(`DELETE FROM sensors WHERE id=?`, id)
	return err
}
