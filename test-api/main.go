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

	erros_channel := make(chan string, len(routes))

	for _, route := range routes {
		wg.Add(1)
		go func(route string, errors_channel chan string) {
			checkAPI(route, errors_channel)
			defer wg.Done()
		}(route, erros_channel)
	}

	wg.Wait()

	for i := 0; i < len(routes); i++ {
		fmt.Println(<-erros_channel)
	}

	elapsed := time.Since(start)
	fmt.Printf("Tooked: %v segundos\n", elapsed.Seconds())
}

func checkAPI(path string, ch chan string) {
	if _, err := http.Get(path); err != nil {
		error_str := fmt.Sprintf("ERROR: %s esta caido \n", err)

		ch <- error_str

		return
	}

	ch <- fmt.Sprintf("%s Esta funcionando perfectamente!", path)

}
