Go Web Crawler
A simple, concurrent web crawler built with Go. This tool starts at a specified base URL and recursively crawls all pages within the same domain, providing a final report of all unique internal pages and the number of times they were linked.

Features
Concurrent Crawling: Utilizes goroutines and channels to fetch multiple pages at once for improved speed.

Domain Scoped: The crawler will only visit pages that are on the same host as the initial base URL, ignoring all external links.

URL Normalization: Cleans up URLs to their canonical form (e.g., https://example.com/path/ becomes example.com/path) to prevent duplicate counting.

Crawl Report: After the crawl is complete, it prints a list of all discovered pages and the number of times each was visited.

Prerequisites
To run this project, you will need:

Go: Version 1.21 or later.

Installation
Clone the repository:

git clone https://github.com/shreyasganesh0/webcrawler-go.git
cd webcrawler-go

Build the executable:

go build -o crawler ./cmd/crawler

This will create an executable file named crawler in your project's root directory.

Usage
Run the crawler from your terminal by providing a single base URL as an argument.

./crawler <BASE_URL>

Example
To crawl a fictional blog and its internal links:

./crawler https://wagslane.dev

Example Output:

starting crawl of: https://wagslane.dev...
crawling https://wagslane.dev
crawling https://wagslane.dev/about/
crawling https://wagslane.dev/posts/
crawling https://wagslane.dev/tags/
...
1 - wagslane.dev/about/
5 - wagslane.dev/tags/
12 - wagslane.dev/
8 - wagslane.dev/posts/

Project Structure
The project is organized with a standard Go layout:

.
├── cmd/
│   └── crawler/
│       └── main.go      # Main application entry point and crawl logic
├── internal/
│   └── url/
│       ├── extract_html.go  # Functions for parsing HTML and extracting links
│       └── normalize_url.go # URL normalization logic
├── go.mod
├── go.sum
└── README.md

cmd/crawler: Contains the main application code. This is the entry point for the executable.

internal/url: A private package containing helper functions for URL manipulation and HTML parsing. Code in internal can only be used by this project.
