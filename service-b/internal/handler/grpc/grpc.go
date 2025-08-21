package grpc

import (
	"context"
	"log"
	"service-b/internal/domain/model"
	"service-b/internal/repository"
	pb "service-b/proto"
	"time"
)

type SensorServer struct {
	pb.UnimplementedSensorServiceServer
	repo repository.SensorRepository
}

func NewSensorServer(repo repository.SensorRepository) *SensorServer {
	return &SensorServer{repo: repo}
}

func (s *SensorServer) SendData(ctx context.Context, in *pb.SensorData) (*pb.Empty, error) {
	t, err := time.Parse(time.RFC3339, in.Timestamp)
	if err != nil {
		log.Printf("timestamp parse error: %v", err)
		return &pb.Empty{}, err
	}
	data := &model.SensorData{
		SensorValue: in.SensorValue,
		SensorType:  in.SensorType,
		ID1:         in.Id1,
		ID2:         int(in.Id2),
		Timestamp:   t,
	}
	err = s.repo.Save(data)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}
