package main

import "fmt"

func main() {
	workerInput := make(chan int)
	workerOutput := make(chan int)
	squareWorker := func() {
		for {
			fmt.Println("masuk")
			num := <-workerInput
			workerOutput <- num * num
		}
	}
	go squareWorker()
	for i := 0; i < 5; i++ {
		var res int
		// begin answer
		workerInput <- i
		res = <-workerOutput
		// end answer
		fmt.Println("res: ", res)
	}

}
