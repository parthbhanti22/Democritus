package store

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/parthbhanti22/democritus/proto"
)

// TaskStatus defines the state of a simulation task.
type TaskStatus int

const (
	StatusQueued    TaskStatus = iota
	StatusRunning              // Dispatched to a worker
	StatusCompleted            // Result received
)

// Task represents a unit of work.
type Task struct {
	ID        string
	Seed      int64
	Status    TaskStatus
	StartedAt time.Time
	WorkerID  string
	Result    *pb.TaskResult
}

// MemoryStore manages tasks in memory with thread safety.
type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]*Task
	queue []*Task // Simple queue for pending tasks
}

// NewMemoryStore initializes the store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: make(map[string]*Task),
		queue: make([]*Task, 0),
	}
}

// AddTask adds a new task to the store.
func (s *MemoryStore) AddTask(t *Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[t.ID] = t
	s.queue = append(s.queue, t)
}

// GetTask retrieves a queued task for a worker.
// It acts as a transaction: finding work and marking it as "Running".
func (s *MemoryStore) GetTask(workerID string) (*Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.queue) == 0 {
		return nil, fmt.Errorf("no work available")
	}

	// Pop from queue
	task := s.queue[0]
	s.queue = s.queue[1:]

	// Update state
	task.Status = StatusRunning
	task.WorkerID = workerID
	task.StartedAt = time.Now()

	return task, nil
}

// CompleteTask marks a task as finished and saves the result.
func (s *MemoryStore) CompleteTask(res *pb.TaskResult) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[res.TaskId]
	if !exists {
		return fmt.Errorf("task %s not found", res.TaskId)
	}

	task.Status = StatusCompleted
	task.Result = res
	// We don't remove it from the map because we might want to query simulations later.
	// But in a real system, we might archive it to a DB.

	return nil
}

// CheckTimeouts is the Reaping logic.
// It finds "Running" tasks that have exceeded the duration and resets them to "Queued".
func (s *MemoryStore) CheckTimeouts(timeout time.Duration) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	reclaimed := 0
	now := time.Now()

	for _, task := range s.tasks {
		if task.Status == StatusRunning {
			if now.Sub(task.StartedAt) > timeout {
				// FAULT TOLERANCE TRIGGERED
				// The worker died or is too slow.
				// Reset status and put back in queue.
				fmt.Printf("REAPER: Reclaiming Task %s from dead worker %s\n", task.ID, task.WorkerID)
				
				task.Status = StatusQueued
				task.WorkerID = ""
				task.StartedAt = time.Time{} // Reset time
				
				// Re-queue
				s.queue = append(s.queue, task) // Add to end of queue
				reclaimed++
			}
		}
	}
	return reclaimed
}
