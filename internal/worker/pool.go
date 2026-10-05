package worker

import (
	"log"
	"time"

	"github.com/AlmatKozhakhmetov/notification-service/internal/domain"
)

type WorkerPool struct {
	taskQueue   chan domain.Notification
	workerCount int
}

func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan domain.Notification, queueSize),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.workerCount; i++ {
		go wp.worker(i)
	}
	log.Printf("Worker pool started with %d workers\n", wp.workerCount)
}

func (wp *WorkerPool) Submit(n domain.Notification) {
	wp.taskQueue <- n
}

func (wp *WorkerPool) worker(id int) {
	for notification := range wp.taskQueue {
		log.Printf("[Worker %d] Processing notification %s (%s) for %s\n",
			id, notification.ID, notification.Type, notification.Recipient)

		time.Sleep(2 * time.Second)

		log.Printf("[Worker %d] SUCCESSFULLY sent notification %s\n", id, notification.ID)
	}
}
