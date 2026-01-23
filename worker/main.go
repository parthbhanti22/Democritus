package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/parthbhanti22/democritus/proto"
	"github.com/parthbhanti22/democritus/worker/simulation"
)

func main() {
	// 1. Connect to the Scheduler (Master)
	address := os.Getenv("SCHEDULER_URL")
	if address == "" {
		address = "localhost:50051"
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDemocritusSchedulerClient(conn)

	// 2. Generate Worker ID
	workerID := uuid.New().String()
	log.Printf("Worker %s starting...", workerID)

	// 3. Register Worker (with Retry)
	var regResp *pb.RegistrationStatus
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		regResp, err = client.RegisterWorker(ctx, &pb.WorkerRegistration{
			WorkerId: workerID,
			Port:     0, // Not listening currently
		})
		cancel()

		if err == nil {
			break
		}
		log.Printf("Failed to register (Master unavailable?): %v. Retrying in 2s...", err)
		time.Sleep(2 * time.Second)
	}
	log.Printf("Registered: %v", regResp.GetSuccess())

	// 4. The Work Loop
	sim := simulation.NewRandomWalk3D()

	for {
		// Ask for work
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		task, err := client.GetTask(ctx, &pb.WorkerID{WorkerId: workerID})
		cancel()

		if err != nil {
			// This likely means no work available or connection error
			// log.Printf("Error getting task: %v", err) // Optional: be quiet if strict no-spam
			time.Sleep(1 * time.Second)
			continue
		}

		if task.GetTaskId() == "" {
			// Empty task?
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("Received Task %s: %d iterations, valid seed: %d", task.GetTaskId(), task.GetIterationCount(), task.GetSeed())

		// Run Simulation
		coords, err := sim.Run(task.GetSeed(), task.GetIterationCount())
		if err != nil {
			log.Printf("Simulation failed: %v", err)
			continue
		}

		// Submit Result
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		_, err = client.SubmitResult(ctx, &pb.TaskResult{
			TaskId:           task.GetTaskId(),
			FinalCoordinates: coords,
		})
		cancel()

		if err != nil {
			log.Printf("Failed to submit result: %v", err)
		} else {
			log.Printf("Task %s completed & submitted.", task.GetTaskId())
		}
	}
}
