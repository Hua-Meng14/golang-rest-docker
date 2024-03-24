# GoLang Project with Docker

This repository contains a GoLang project that is set up to run inside a Docker container. The project utilizes Docker Compose for container orchestration and includes a simple setup to get you started with development.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker installed on your machine. You can download Docker from [here](https://www.docker.com/get-started).
- Basic understanding of Docker and Docker Compose.
- GoLang installed on your machine if you plan to develop Go applications locally.

## Important Commands

Below are nessecary command lines to follow:

1. Docker Compose:

   ```bash
   docker compose build
   
   docker compose up
   
   docker compose down
   
   docker compose run --service-ports web bash
   ```
2. Go Download libraries:

   ```bash
   go mod init <github_repo_url>
   
   go get github.com/gofiber/fiber/v2
   
   go get gorm.io/gorm
   
   go get gorm.io/driver/postgres
   ```

3. Bind container to hostmachine

   ```bash
   go run cmd/main.go -b 0.0.0.0
   ```

