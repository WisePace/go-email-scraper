package main

import (
	spinner "email-scraper/pkg/animation"
	"email-scraper/pkg/log"
	"email-scraper/pkg/scraper"
	logger "log"
	"os"
	"time"
)

func main() {
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
	s := spinner.New([]rune{'|', '/', '-', '\\'}, 100*time.Millisecond)
	s.Start()

	// Read existing emails into a set
	s.Suffix(" Reading existing emails...")
	existingEmails, err := scraper.ReadExistingEmails("email_list.txt")
	if err != nil {
		s.Stop()
		errorLogger.Printf("Error reading existing emails: %v", err)
		logger.Printf("Error reading existing emails: %v\n", err)
		return
	}
	s.Suffix(" Successfully read existing emails.")
	logger.Println("Successfully read existing emails.")

	// Reopen email list file in append mode
	s.Suffix(" Opening email list file in append mode...")
	emailListFile, err := os.OpenFile("email_list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		s.Stop()
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
	s.Suffix(" Reading domains from file...")
	domains, err := scraper.ReadDomains(file)
	if err != nil {
		s.Stop()
		errorLogger.Printf("Error reading domains: %v", err)
		return
	}
	logger.Println("Successfully read domains from file.")

	// Find emails for all domains
	s.Suffix(" Finding emails for all domains...")
	foundEmails, scannedDomains := scraper.FindEmails(domains, emailListFile, errorLogger, existingEmails)

	s.Stop()
	logger.Printf("Scanned domains: %d, Found emails: %d\n", scannedDomains, foundEmails)
	logger.Printf("Scanned domains: %d, Found emails: %d\n", scannedDomains, foundEmails)
}
