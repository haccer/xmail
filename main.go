package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/haccer/available"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	uniqueDomains := make(map[string]bool)

	var wg sync.WaitGroup
	workerCount := 100
	jobs := make(chan string, workerCount)

	for w := 1; w <= workerCount; w++ {
		go func() {
			for domain := range jobs {
				available := available.Domain(domain)
				if available {
					fmt.Printf("[+] Available Domain: %s\n", domain)
				}
				wg.Done()
			}
		}()
	}

	for scanner.Scan() {
		email := scanner.Text()
		at := strings.LastIndex(email, "@")
		if at >= 0 {
			domain := email[at+1:]
			uniqueDomains[domain] = true
		}
	}

	for domain := range uniqueDomains {
		wg.Add(1)
		jobs <- domain
	}
	close(jobs)

	wg.Wait()
}
