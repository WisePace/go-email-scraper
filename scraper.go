package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func FindEmails(domains []string, emailListFile *os.File, logger *log.Logger, existingEmails map[string]struct{}) {
	var wg sync.WaitGroup

	threadCount := os.Getenv("THREAD_COUNT")
	if threadCount == "" {
		threadCount = "10"
	}

	threads, err := strconv.Atoi(threadCount)
	if err != nil {
		logger.Printf("Error parsing THREAD_COUNT: %v", err)
	}

	buffer := make(chan struct{}, threads) // Buffered channel to control concurrency
	for _, domain := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			buffer <- struct{}{}        // Acquire the token
			defer func() { <-buffer }() // Release the token

			url := "https://" + domain
			fmt.Printf("Scraping emails from %s\n", url)
			emails, err := scrapeEmailsFromDomain(url)
			if err != nil {
				logger.Printf("Error scraping %s: %v", url, err)
				return
			}

			for _, email := range emails {
				if _, exists := existingEmails[email]; !exists {
					logger.Printf("Found email: %s", email)
					fmt.Printf("Found an email: %s\n", email)
					emailListFile.WriteString(fmt.Sprintf("%s\n", email))
					existingEmails[email] = struct{}{}
				}
			}
		}(domain)
	}

	wg.Wait()
}

func scrapeEmailsFromDomain(url string) ([]string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emailSet := make(map[string]struct{})

	for {
		nextToken := tokenizer.Next()
		if nextToken == html.ErrorToken {
			break
		}
		if nextToken == html.TextToken {
			text := string(tokenizer.Text())
			matches := emailRegex.FindAllString(text, -1)
			for _, email := range matches {
				emailSet[email] = struct{}{}
			}
		}
	}

	var emails []string
	for email := range emailSet {
		emails = append(emails, email)
	}
	return emails, nil
}
