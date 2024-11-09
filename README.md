# WiseScraper

WiseScraper is a simple email scraper written in Go. It reads a list of domains from a file, scrapes emails from the websites, and writes the found emails to an output file. 

The scraper supports concurrent scraping with a configurable thread count.

## Features

- Scrapes emails from a list of domains
- Supports concurrent scraping with configurable thread count
- Logs errors and found emails
- Reads existing emails to avoid duplicates

## Prerequisites

- Go 1.23 or later

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/emailscraper.git
    cd emailscraper
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Email Scraper

1. Prepare your `domains.txt` file with a list of domains to scrape, one per line.

2. Run the scraper:
    ```sh
    go run main.go
    ```

3. The scraper will read the domains from `domains.txt`, scrape emails, and append them to `email_list.txt`. Logs will be written to `app.log` and `error.log`.

## Configuration

- Set the `THREAD_COUNT` environment variable to control the number of concurrent scraping threads (default is 10).
- Set the `DOMAINS_FILE` environment variable to specify a different domains file (default is `domains.txt`).

Example:
```sh
THREAD_COUNT=20 DOMAINS_FILE=domains.txt go run main.go
```

## File Structure

- `main.go`: Entry point for the email scraper.
- `scraper.go`: Contains the scraping logic.
- `domains.txt`: List of domains to scrape.
- `email_list.txt`: Output file for found emails.
- `app.log`: Log file for general logs.
- `error.log`: Log file for error logs.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.