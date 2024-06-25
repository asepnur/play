package main

import "fmt"

func main() {
	executed := false
	terminate(func() {
		c := make(chan int)
		f := func() { <-c }
		//beginanswer
		go f()
		//endanswer f()
		c <- 2

		executed = true
	})
	fmt.Println(executed)

}

func terminate(f func()) {
	done := make(chan interface{})
	go func() {
		f()
		close(done)
	}()
}
