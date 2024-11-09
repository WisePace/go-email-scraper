package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open log files
	appLogs, _ := OpenLogFile("app.log")
	defer appLogs.Close()
	log.SetOutput(appLogs)

	errorLogs, errorLogger := OpenLogFile("error.log")
	defer errorLogs.Close()

	// Read existing emails into a set
	existingEmails, err := readExistingEmails("email_list.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Reopen email list file in append mode
	emailListFile, err := os.OpenFile("email_list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening email list file: %v", err)
	}
	defer emailListFile.Close()

	file := os.Getenv("DOMAINS_FILE")
	if file == "" {
		file = "domains.txt"
	}

	// Read domains from domains.txt
	domains, err := readDomains(file)
	if err != nil {
		errorLogger.Fatalf("%v", err)
	}

	// Find emails for all domains
	FindEmails(domains, emailListFile, errorLogger, existingEmails)
}

func readExistingEmails(fileName string) (map[string]struct{}, error) {
	existingEmails := make(map[string]struct{})

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening email list file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		existingEmails[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading email list file: %v", err)
	}

	return existingEmails, nil
}

func readDomains(fileName string) ([]string, error) {
	var domains []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening domains file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading domains file: %v", err)
	}

	return domains, nil
}
