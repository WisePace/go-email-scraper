package scraper

import (
	"bufio"
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

func ReadExistingEmails(fileName string) (map[string]struct{}, error) {
	existingEmails := make(map[string]struct{})

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Printf("%s does not exist. Proceeding without reading existing emails.\n", fileName)
		return existingEmails, nil
	}

	log.Printf("Opening email list file: %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening email list file: %v", err)
	}
	defer file.Close()

	log.Println("Reading emails from file...")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		email := scanner.Text()
		existingEmails[email] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading email list file: %v", err)
	}

	log.Println("Finished reading emails from file.")
	return existingEmails, nil
}

func ReadDomains(fileName string) ([]string, error) {
	var domains []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening domains file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading domains file: %v", err)
	}

	return domains, nil
}

func FindEmails(domains []string, emailListFile *os.File, logger *log.Logger, existingEmails map[string]struct{}) int {
	var wg sync.WaitGroup
	var scannedDomains int

	threadCount := os.Getenv("THREAD_COUNT")
	if threadCount == "" {
		threadCount = "100" // Increase the default thread count
	}

	threads, err := strconv.Atoi(threadCount)
	if err != nil {
		logger.Printf("Error parsing THREAD_COUNT: %v", err)
		threads = 100
	}

	// Use a buffered channel to limit the number of concurrent goroutines
	buffer := make(chan struct{}, threads)
	for _, domain := range domains {
		wg.Add(1)
		scannedDomains++
		go func(domain string) {
			defer wg.Done()
			buffer <- struct{}{}        // Acquire the token
			defer func() { <-buffer }() // Release the token

			url := "https://" + domain
			emails, err := scrapeEmailsFromDomain(url)
			if err != nil {
				logger.Printf("Error scraping %s: %v", url, err)
				return
			}

			for _, email := range emails {
				if _, exists := existingEmails[email]; !exists {
					logger.Printf("Found email: %s", email)
					if _, err := emailListFile.WriteString(fmt.Sprintf("%s\n", email)); err != nil {
						logger.Printf("Error writing email to file: %v", err)
					}
					existingEmails[email] = struct{}{}
				}
			}
		}(domain)
	}

	wg.Wait()
	return scannedDomains
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

	emails := make([]string, 0, len(emailSet))
	for email := range emailSet {
		emails = append(emails, email)
	}
	return emails, nil
}
