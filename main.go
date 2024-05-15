package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/haccer/available"
)

var wordlistFile string

func init() {
	flag.StringVar(&wordlistFile, "w", "", "Wordlist file")
}

func main() {
	flag.Parse()

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
					fmt.Println(domain)
				}
				wg.Done()
			}
		}()
	}

	if wordlistFile != "" {
		file, err := os.Open(wordlistFile)
		if err != nil {
			fmt.Println("Error opening wordlist file:", err)
			return
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			email := fileScanner.Text()
			at := strings.LastIndex(email, "@")
			if at >= 0 {
				domain := email[at+1:]
				uniqueDomains[domain] = true
			}
		}
	} else {
		for scanner.Scan() {
			email := scanner.Text()
			at := strings.LastIndex(email, "@")
			if at >= 0 {
				domain := email[at+1:]
				uniqueDomains[domain] = true
			}
		}
	}

	for domain := range uniqueDomains {
		wg.Add(1)
		jobs <- domain
	}
	close(jobs)

	wg.Wait()
}
