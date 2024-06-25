package main

import (
	"fmt"
	"log"

	"os"
	"os/signal"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

// Make an enqueuer with a particular namespace
var enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)

func main() {
	worker()
	return
	// Enqueue a job named "send_email" with the specified parameters.
	_, err := enqueuer.EnqueueUnique("send_email", work.Q{"address": "5@example.com", "subject": "hello world", "customer_id": 5})
	// Enqueue a job named "send_email" with the specified parameters.
	_, err = enqueuer.EnqueueUnique("send_email", work.Q{"address": "6@example.com", "subject": "hello world", "customer_id": 6})
	_, err = enqueuer.EnqueueUnique("send_email", work.Q{"address": "7@example.com", "subject": "hello world", "customer_id": 7})
	_, err = enqueuer.EnqueueUnique("send_email", work.Q{"address": "7@example.com", "subject": "hello world", "customer_id": 7})
	_, err = enqueuer.Enqueue("send_email", work.Q{"address": "7@example.com", "subject": "hello world", "customer_id": 7})
	_, err = enqueuer.Enqueue("send_email", work.Q{"address": "8@example.com", "subject": "hello world", "customer_id": 8})
	if err != nil {
		log.Fatal(err)
	}
}

// Make a redis pool
type Context struct {
	customerID int64
}

func worker() {
	// Make a new pool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "my_app_namespace" is the Redis namespace
	// redisPool is a Redis pool
	pool := work.NewWorkerPool(Context{}, 10, "my_app_namespace", redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).FindCustomer)

	// Map the name of jobs to handler functions
	pool.JobWithOptions("send_email", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).SendEmail)

	// Customize options:
	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a customer_id param, set it in the context for future middleware and handlers to use.
	if _, ok := job.Args["customer_id"]; ok {
		c.customerID = job.ArgInt64("customer_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}

	return next()
}

func (c *Context) SendEmail(job *work.Job) error {
	// Extract arguments:
	addr := job.ArgString("address")
	subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}
	log.Println(subject)
	log.Println(addr)
	// Go ahead and send the email...
	// sendEmailTo(addr, subject)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}
