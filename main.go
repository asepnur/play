package main

import (
	"fmt"
	"log"
	"time"
)

const asep = ""

func main() {
	var t time.Time
	log.Println("time 0: ", t.UTC().Format(time.RFC3339))
	lastDate := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	log.Println("last time: ", lastDate.Format(time.RFC3339))
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("time: ", time.Now().In(loc))
}
