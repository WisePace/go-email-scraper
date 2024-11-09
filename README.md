# WiseScraper

WiseScraper is a simple email scraper written in Go. It reads a list of domains from a file, scrapes emails from the websites, and writes the found emails to an output file. 

The scraper supports concurrent scraping with a configurable thread count.

## Features

- Scrapes emails from a list of domains
- Supports concurrent scraping with configurable thread count
- Logs errors and found emails
- Reads existing emails to avoid duplicates

## Prerequisites

- Go 1.23 or later (for building from source)

## Installation

### From Releases

1. Download the appropriate binary for your platform:
    ```sh
    wget https://github.com/WisePace/pace-scraper/releases/download/v1.0.0/email-scraper-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m)
    ```

2. Make the binary executable:
    ```sh
    chmod +x email-scraper-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m)
    ```

3. Move the binary to a directory in your `PATH` (e.g., `/usr/local/bin`):
    ```sh
    sudo mv email-scraper-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m) /usr/local/bin/email-scraper
    ```

### From Source

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/emailscraper.git
    cd emailscraper
    ```

2. Install Go dependencies:
    ```sh
    go mod tidy
    ```

3. Build the binary:
    ```sh
    go build -o email-scraper main.go
    ```

## Usage

### Email Scraper

1. Prepare your `domains.txt` file with a list of domains to scrape, one per line.

2. Run the scraper:
    ```sh
    email-scraper
    ```

The scraper will read the domains from `domains.txt`, scrape emails, and append them to `email_list.txt`. 

Logs will be written to `app.log` and `error.log`.

## Configuration

- Set the `THREAD_COUNT` environment variable to control the number of concurrent scraping threads (default is 10).
- Set the `DOMAINS_FILE` environment variable to specify a different domains file (default is `domains.txt`).

Example:
```sh
THREAD_COUNT=20 DOMAINS_FILE=domains.txt email-scraper
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

- [Christofher](https://github.com/KostLinux)