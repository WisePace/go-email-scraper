package main

import (
	"bufio"
	"email-scraper/pkg/log"
	"email-scraper/pkg/scraper"
	"email-scraper/pkg/spinner"
	"fmt"
	logger "log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	logger.Println("Starting the email scraper...")

	// Open log files
	appLogs, err := log.OpenFile("app.log")
	if err != nil {
		logger.Printf("Error opening app log file: %v\n", err)
		return
	}
	defer appLogs.Close()
	logger.SetOutput(appLogs)
	logger.Println("App log file opened successfully.")

	errorLogs, err := log.OpenFile("error.log")
	if err != nil {
		logger.Printf("Error opening error log file: %v\n", err)
		return
	}
	defer errorLogs.Close()
	errorLogger := logger.New(errorLogs, "ERROR: ", logger.LstdFlags)
	errorLogger.Println("Error log file opened successfully.")

	// Initialize spinner
	spin := spinner.New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	spin.Start()

	// Read existing emails into a set
	spin.Suffix(" Reading existing emails...")
	existingEmails, err := scraper.ReadExistingEmails("email_list.txt")
	if err != nil {
		spin.Stop()
		errorLogger.Printf("Error reading existing emails: %v", err)
		logger.Printf("Error reading existing emails: %v\n", err)
		return
	}
	spin.Suffix(" Successfully read existing emails.")
	logger.Println("Successfully read existing emails.")

	// Reopen email list file in append mode
	spin.Suffix(" Opening email list file in append mode...")
	emailListFile, err := os.OpenFile("email_list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		spin.Stop()
		errorLogger.Printf("Error opening email list file: %v", err)
		return
	}
	defer emailListFile.Close()
	logger.Println("Email list file opened in append mode.")

	file := os.Getenv("DOMAINS_FILE")
	if file == "" {
		file = "domains.txt"
	}

	// Read domains from domains.txt
	spin.Suffix(" Reading domains from file...")
	domains, err := scraper.ReadDomains(file)
	if err != nil {
		spin.Stop()
		errorLogger.Printf("Error reading domains: %v", err)
		return
	}
	logger.Println("Successfully read domains from file.")

	// Find emails for all domains
	spin.Suffix(" Finding emails for all domains...")
	scannedDomains := scraper.FindEmails(domains, emailListFile, errorLogger, existingEmails)

	spin.Stop()
	duration := time.Since(startTime)

	// Count emails from file
	emailCount, err := countLines("email_list.txt")
	if err != nil {
		logger.Printf("Error counting emails: %v\n", err)
		fmt.Printf("Error counting emails: %v\n", err)
		return
	}

	logger.Printf("Scanned domains: %d, Found emails: %d, Time consumed: %s\n", scannedDomains, emailCount, duration)
	fmt.Printf("Scanned domains: %d, Found emails: %d, Time consumed: %s\n", scannedDomains, emailCount, duration)
}

func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return count, nil
}
