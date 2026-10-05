package main

import (
	"log"
	"net"

	
import (
    "github.com/AlmatKozhakhmetov/-notification-service-/internal/handler"
    "github.com/AlmatKozhakhmetov/-notification-service-/internal/worker"
)

func main() {
	// 1. Запускаем Worker Pool
	pool := worker.NewWorkerPool(3, 100)
	pool.Start()

	// 2. Слушаем TCP-порт 50051 (стандартный порт для gRPC)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// 3. Регистрируем gRPC-сервер
	grpcServer := grpc.NewServer()
	grpcHandler := handler.NewGRPCHandler(pool)

	pb.RegisterNotificationServiceServer(grpcServer, grpcHandler)

	log.Println("gRPC Notification service running on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
