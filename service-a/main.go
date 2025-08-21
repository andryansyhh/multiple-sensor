package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	rpc "service-a/internal/handler/grpc"
	handler "service-a/internal/handler/http"
	"service-a/internal/usecase"
	pb "service-a/proto"
)

func main() {
	// Sensor type
	sensorType := os.Getenv("SENSOR_TYPE")
	if sensorType == "" {
		sensorType = "default"
	}

	// gRPC target
	grpcTarget := os.Getenv("GRPC_TARGET")
	if grpcTarget == "" {
		grpcTarget = "localhost:50051"
	}
	conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	client := pb.NewSensorServiceClient(conn)

	// init generator + sender
	generator := usecase.NewDataGenerator(sensorType)
	sender := rpc.NewDataSender(client)

	// goroutine buat kirim data periodik
	go func() {
		for {
			data := generator.Generate()
			sender.Send(data)
			time.Sleep(handler.GetFrequency())
		}
	}()

	// REST API
	e := echo.New()
	handler.SetupRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting Service A with SENSOR_TYPE=%s on PORT=%s", sensorType, port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
