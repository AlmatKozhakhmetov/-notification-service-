package handler

import (
	pb "github.com/AlmatKozhakhmetov/-notification-service-/api/proto"
	"github.com/AlmatKozhakhmetov/-notification-service-/internal/worker"
)

// GRPCHandler реализует gRPC-сервер для уведомлений
type GRPCHandler struct {
	pb.UnimplementedNotificationServiceServer
	workerPool *worker.Pool
}

// NewGRPCHandler — конструктор для создания хендлера
func NewGRPCHandler(wp *worker.Pool) *GRPCHandler {
	return &GRPCHandler{
		workerPool: wp,
	}
}
