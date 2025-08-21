package grpc

import (
	"context"
	"log"

	"service-a/internal/domain/model"
	pb "service-a/proto"
)

type DataSender struct {
	client pb.SensorServiceClient
}

func NewDataSender(client pb.SensorServiceClient) *DataSender {
	return &DataSender{client: client}
}

func (s *DataSender) Send(data *model.SensorData) {
	_, err := s.client.SendData(context.Background(), &pb.SensorData{
		SensorValue: data.SensorValue,
		SensorType:  data.SensorType,
		Id1:         data.ID1,
		Id2:         int32(data.ID2),
		Timestamp:   data.Timestamp.Format("2006-01-02T15:04:05Z07:00"),
	})
	if err != nil {
		log.Printf("failed to send data: %v", err)
	}
}
