package main

import (
	"sync"
	"time"
)

func processMessages(messages []string) []string {
	var pm []string
	var wg sync.WaitGroup
	for _, m := range messages {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pm = append(pm, process(m))
		}()
	}
	wg.Wait()
	return pm
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
