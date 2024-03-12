package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	start := time.Now()

	routes := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet",
		"https://graph.microsoft.com",
	}

	for _, route := range routes {
		wg.Add(1)
		go func(route string) {
			checkAPI(route)
			defer wg.Done()
		}(route)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Tooked: %v segundos\n", elapsed.Seconds())
}

func checkAPI(path string) {
	if _, err := http.Get(path); err != nil {
		fmt.Printf("ERROR: %s esta caido \n", err)
		return
	}
}
