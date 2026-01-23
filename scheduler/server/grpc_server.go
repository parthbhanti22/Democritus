package server

import (
	"context"
	"fmt"
	"log"

	pb "github.com/parthbhanti22/democritus/proto"
	"github.com/parthbhanti22/democritus/scheduler/store"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	tasksCompleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tasks_completed_total",
		Help: "The total number of processed tasks",
	})
)

// GrpcServer implements the DemocritusSchedulerServer interface.
type GrpcServer struct {
	pb.UnimplementedDemocritusSchedulerServer
	Store *store.MemoryStore
}

func NewGrpcServer(s *store.MemoryStore) *GrpcServer {
	return &GrpcServer{Store: s}
}

// RegisterWorker logs a new connection.
func (s *GrpcServer) RegisterWorker(ctx context.Context, req *pb.WorkerRegistration) (*pb.RegistrationStatus, error) {
	log.Printf("Worker Registered: %s", req.WorkerId)
	return &pb.RegistrationStatus{Success: true, Message: "Welcome to the Grid"}, nil
}

// GetTask assigns a task to a requesting worker.
func (s *GrpcServer) GetTask(ctx context.Context, req *pb.WorkerID) (*pb.TaskPayload, error) {
	task, err := s.Store.GetTask(req.WorkerId)
	if err != nil {
		// No work available is not a fatal error, purely an "empty" state.
		// Return empty payload (Worker will sleep)
		return &pb.TaskPayload{}, nil 
	}

	log.Printf("Dispatching Task %s to Worker %s", task.ID, req.WorkerId)

	return &pb.TaskPayload{
		TaskId:         task.ID,
		Seed:           task.Seed,
		IterationCount: 1_000_000, // Fixed size for MVP, could be dynamic
		SimulationType: "random_walk_3d",
	}, nil
}

// SubmitResult processes the computation result.
func (s *GrpcServer) SubmitResult(ctx context.Context, req *pb.TaskResult) (*pb.Ack, error) {
	log.Printf("Received Result for Task %s (Coords: %.2f, %.2f, %.2f)", 
		req.TaskId, 
		req.FinalCoordinates[0], 
		req.FinalCoordinates[1],
		req.FinalCoordinates[2])

	err := s.Store.CompleteTask(req)
	if err != nil {
		return &pb.Ack{Received: false}, fmt.Errorf("failed to save result: %v", err)
	}

	tasksCompleted.Inc() // Metric +1

	return &pb.Ack{Received: true}, nil
}
