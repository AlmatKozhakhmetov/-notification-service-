package main

import (
	"log"
	"net/http"

	"github.com/AlmatKozhakhmetov/notification-service/internal/handler"
	"github.com/AlmatKozhakhmetov/notification-service/internal/worker"
)

func main() {
	pool := worker.NewWorkerPool(3, 100)
	pool.Start()

	notifHandler := handler.NewNotificationHandler(pool)

	http.HandleFunc("/api/v1/notifications", notifHandler.Send)

	log.Println("Notification service running on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
