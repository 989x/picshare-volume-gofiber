# picshare-volume-gofiber

## Project Structure

```
picshare-volume-gofiber/
├── cmd/
│   └── main.go
├── static/
│   └── images/
│       ├── news/
│       └── articles/
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

## Initialize a Go Module

Run the following commands to initialize the Go module and install dependencies:

```bash
go mod init picshare-volume-gofiber
go get github.com/gofiber/fiber/v2
```

## Build & Run with Docker Compose

To build and run the application using Docker Compose, run:

```bash
docker-compose up --build
```

This command will:
1. Build the Docker image from the Dockerfile.
2. Start the container with the required configuration.

To stop the application and remove the containers:

```bash
docker-compose down
```
