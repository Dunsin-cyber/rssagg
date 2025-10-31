# RSS Server

A simple RSS feed server built with Go, allowing users to aggregate and serve RSS feeds.

## Description

This project is a RSS server implementation following the tutorial by Wagslane on YouTube. It provides functionality to fetch, parse and serve RSS feeds through a REST API.

## Features

- RSS feed aggregation
- REST API endpoints
- Feed parsing and validation
- [Add more features as implemented]

## Prerequisites

- Go 1.25.3
- sqlc
- gooose

## Installation

1. Clone the repository

```bash
git clone https://github.com/Dunsin-cyber/go_server.git
cd go_server
```

2. Install dependencies

```bash
go mod download
```
3. Set up Environment Variables
Ensure you use a valid postgres database connection url
```bash 
cp .env.example .env
```

4. Run the server

```bash
go run main.go
```

## API Endpoints

[Document your API endpoints here]

## Tutorial Reference

This project was built following the tutorial series available at:

https://www.youtube.com/watch?v=un6ZyFkqFKo&list=PPSV

## Contributing

Feel free to submit issues and enhancement requests.

## License

[MIT]

## Author

[Dunsin/Dusin-cyber]