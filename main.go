package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/haccer/available"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		email := scanner.Text()
		at := strings.LastIndex(email, "@")
		if at >= 0 {
			domain := email[at+1:]

			available := available.Domain(domain)
			if available {
				fmt.Printf("[+] Available Domain: %s | Email: %s\n", domain, email)
			}
		}
	}
}
