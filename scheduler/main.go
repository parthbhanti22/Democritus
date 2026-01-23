package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/parthbhanti22/democritus/proto"
	"github.com/parthbhanti22/democritus/scheduler/server"
	"github.com/parthbhanti22/democritus/scheduler/store"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	port        = ":50051"
	taskTimeout = 30 * time.Second
	totalTasks  = 10_000
)

func main() {
	// 1. Initialize Store
	ms := store.NewMemoryStore()

	// 2. Populate with Dummy Work (Seeds 1 to 10000)
	// This simulates a large physics simulation job being submitted.
	log.Println("Initializing Job Queue...")
	for i := 1; i <= totalTasks; i++ {
		ms.AddTask(&store.Task{
			ID:     fmt.Sprintf("task-%d", i),
			Seed:   int64(i),
			Status: store.StatusQueued,
		})
	}
	log.Printf("Queue populated with %d tasks.", totalTasks)

	// 3. Start the Reaper (The Fault Tolerance Guardian)
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			reclaimed := ms.CheckTimeouts(taskTimeout)
			if reclaimed > 0 {
				log.Printf("WARNING: Reclaimed %d timed-out tasks!", reclaimed)
			}
		}
	}()

	// 4. Start Metrics Server (Side Door)
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Metrics server listening on :2112")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	// 5. Start gRPC Service
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := server.NewGrpcServer(ms)
	pb.RegisterDemocritusSchedulerServer(s, srv)

	log.Printf("Scheduler listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
