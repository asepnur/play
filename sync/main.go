package main

import "sync"

func main() {
	var passwords []string
	errors := make(chan error, len(passwords))
	results := make(chan string, len(passwords))
	var wg sync.WaitGroup
	wg.Add(len(passwords))
	for _, password := range passwords {
		go func(pwd string) {
			defer wg.Done()
			result, err := createHashedPassword(pwd)
			if err != nil {
				errors <- err
				return
			}

			results <- result
		}(password)
	}

	// do something else
	for range passwords {
		select {
		case err := <-errors:
			println(err)
		case res := <-results:
			println(res)
		}
	}
	wg.Wait()
	// When you are ready to read from goroutine do this:

}
func createHashedPassword(s string) (string, error) {
	return "response", nil
}
