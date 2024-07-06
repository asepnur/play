package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func runEnqueuer() {
	redisConnOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}

	client := asynq.NewClient(redisConnOpt)
	defer client.Close()

	task1 := asynq.NewTask("email:send", []byte(`{"to":"user@example.com","subject":"Welcome"}`))
	task2 := asynq.NewTask("email:send", []byte(`{"to":"user2@example.com","subject":"Welcome"}`))
	var tasks []*asynq.Task
	tasks = append(tasks, task1, task2)
	for _, t := range tasks {
		info, err := client.Enqueue(t)
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}

		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}
}

func runWorker() {
	redisConnOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:send", HandleEmailSendTask)

	srv := asynq.NewServer(
		redisConnOpt,
		asynq.Config{
			Concurrency: 1,
		},
	)
	go logQueueStatus()

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func HandleEmailSendTask(ctx context.Context, t *asynq.Task) error {
	var payload struct {
		To      string
		Subject string
	}

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	log.Printf("Sending email to %s with subject %q", payload.To, payload.Subject)
	// Insert email sending logic here...
	time.Sleep(1 * time.Second)
	log.Println("Sleeping in 1s")
	time.Sleep(30 * time.Second)
	log.Println("Sleeping in 3s")
	time.Sleep(5 * time.Second)
	log.Println("Sleeping in 5s")

	return nil
}

func logQueueStatus() {
	redisConnOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	inspector := asynq.NewInspector(redisConnOpt)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		queues, err := inspector.Queues()
		if err != nil {
			log.Printf("Error inspecting queues: %v", err)
			continue
		}

		allEmpty := true
		for _, q := range queues {
			stats, err := inspector.GetQueueInfo(q)
			if err != nil {
				log.Printf("Error getting queue info for %s: %v", q, err)
				continue
			}
			log.Printf("Queue: %s, Pending tasks: %d, Running: %d, Latency latest task: %+v", stats.Queue, stats.Pending, stats.Active, stats.Latency)
			if stats.Pending > 0 {
				allEmpty = false
			}
		}

		if allEmpty {
			log.Println("All queues are empty")
		}
	}
}

func main() {
	mode := flag.String("mode", "worker", "Mode to run: worker or enqueuer")
	flag.Parse()

	switch *mode {
	case "enqueuer":
		runEnqueuer()
	case "worker":
		runWorker()
	default:
		log.Fatalf("Unknown mode: %s", *mode)
	}
}
