// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type URLConfig struct {
// 	Token    string
// 	Timeout  time.Duration
// 	MaxCount int
// }

// type URLResponse struct {
// 	Result string
// 	URL    string
// 	Error  error
// }

// func FeatchUrl(ctx context.Context, cfg URLConfig, urls []string) []URLResponse {
// 	var wg sync.WaitGroup
// 	result := make(chan URLResponse, len(urls))
// 	semaphore := make(chan struct{}, cfg.MaxCount)

// 	client := &http.Client{
// 		Timeout: cfg.Timeout,
// 	}

// 	for _, url := range urls {
// 		wg.Add(1)

// 		go func(u string) {
// 			defer wg.Done()
// 			semaphore <- struct{}{}

// 			defer func() {
// 				<-semaphore
// 			}()

// 			req, err := http.NewRequestWithContext("GET", url, nil)
// 			if err != nil {
// 				fmt.Println("Error to create request:", err)
// 				return
// 			}
// 			re

// 		}(url)
// 		go func() {
// 			wg.Wait()
// 			close(result)

// 		}()
// 	}
// }

// func main() {

// }

package main
