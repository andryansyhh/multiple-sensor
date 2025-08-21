package dto

// queryParams untuk validasi query parameter
type QueryParams struct {
	ID1    string `query:"id1" validate:"omitempty,alphanum"`
	ID2    string `query:"id2" validate:"omitempty,number"`
	Start  string `query:"start" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	End    string `query:"end" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	Limit  string `query:"limit" validate:"omitempty,number,gte=1,lte=100"`
	Offset string `query:"offset" validate:"omitempty,number,gte=0"`
}

// updateRequest untuk validasi body PUT /data
type UpdateRequest struct {
	SensorValue float64 `json:"sensor_value" validate:"required"`
	SensorType  string  `json:"sensor_type" validate:"required,alphanum"`
}
