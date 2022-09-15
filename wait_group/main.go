package main

import (
	"log"
	"os"
	"sync"
)

func main() {

	logger := log.New(os.Stdout, "", 0)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		logger.Println("From 1")
	}()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Println("From 2")
	}()

	// wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Println("From 3")
	}()

	logger.Println("From main")

	wg.Wait()
}
