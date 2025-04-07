# picshare-volume-gofiber

## Initialize a Go module

```bash
go mod init picshare-volume-gofiber

go get github.com/gofiber/fiber/v2
```

## Build & Run Docker

```bash
docker build -t picshare-volume-gofiber .

docker run -p 8080:8080 picshare-volume-gofiber
```
