package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func CountOccurrencesGo(requestURL string) (count int) {
	resp, err := http.Get(requestURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	return strings.Count(string(body), "Go")
}

type job struct {
	url     string
	counter int
}

func main() {
	const k int = 5
	parallelJobsCh := make(chan struct{}, k)
	resultCh := make(chan job)
	wg := sync.WaitGroup{}

	var total int

	go func(resultCh chan job, wg *sync.WaitGroup) {
		for job := range resultCh {
			fmt.Printf("Count for %s: %d \n", job.url, job.counter)
			total += job.counter
			wg.Done()
		}
	}(resultCh, &wg)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parallelJobsCh <- struct{}{}
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup, parallelJobsCh chan struct{}, result chan job) {
			resultCh <- job{url: url, counter: CountOccurrencesGo(url)}
			<-parallelJobsCh
		}(scanner.Text(), &wg, parallelJobsCh, resultCh)
	}

	wg.Wait()
	close(resultCh)

	fmt.Printf("Total: %d \n", total)
}
