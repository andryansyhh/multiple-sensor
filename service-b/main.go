package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"

	rpc "service-b/internal/handler/grpc"
	"service-b/internal/handler/http"
	"service-b/internal/repository"
	"service-b/internal/usecase"
	pb "service-b/proto"
)

func main() {
	// MySQL
	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = "user:password@tcp(mysql:3306)/sensor_db?parseTime=true"
	}
	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// init repo, usecase, handler
	repo := repository.NewSensorRepository(db)
	usecase := usecase.NewSensorUsecase(repo)
	httpCtrl := http.NewSensorController(usecase)
	grpcServer := rpc.NewSensorServer(usecase)

	// gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSensorServiceServer(s, grpcServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// HTTP server
	e := echo.New()
	http.SetupRoutes(e, httpCtrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	log.Printf("Starting Service B on PORT=%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
