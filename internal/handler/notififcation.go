package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlmatKozhakhmetov/notification-service/internal/domain"
	"github.com/AlmatKozhakhmetov/notification-service/internal/worker"
)

type NotificationHandler struct {
	pool *worker.WorkerPool
}

func NewNotificationHandler(pool *worker.WorkerPool) *NotificationHandler {
	return &NotificationHandler{pool: pool}
}

type SendRequest struct {
	Type      domain.NotificationType `json:"type"`
	Recipient string                  `json:"recipient"`
	Message   string                  `json:"message"`
}

func (h *NotificationHandler) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	notification := domain.Notification{
		ID:        fmt.Sprintf("notif_%d", time.Now().UnixNano()),
		Type:      req.Type,
		Recipient: req.Recipient,
		Message:   req.Message,
		CreatedAt: time.Now(),
	}

	h.pool.Submit(notification)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "queued",
		"task_id": notification.ID,
	})
}
