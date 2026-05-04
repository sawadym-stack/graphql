# Social Post Board (GraphQL + Go)

This is a complete implementation of a GraphQL API built with Go, gqlgen, and GORM (PostgreSQL). It supports full CRUD operations, real-time GraphQL subscriptions, and JWT-based authentication. The application is Dockerized and includes a CI/CD pipeline for automated deployment.

## Entities
- **User**: Can sign up, log in, create posts, and comment on posts.
- **Post**: Users can create, update, delete posts. A post contains multiple comments.
- **Comment**: Users can comment on posts. Real-time updates are sent to clients via subscriptions.

## Prerequisites
- Go 1.21+
- PostgreSQL
- Docker

## Setup Instructions

### 1. Environment Variables
Copy the example environment file and update it with your database credentials:
```bash
cp .env.example .env
```

### 2. Install Dependencies
```bash
go mod tidy
```
*(Note: If you encounter issues with `godotenv`, ensure you run `go get github.com/joho/godotenv`)*

### 3. Running Locally
Start the server:
```bash
go run server.go
```
Navigate to `http://localhost:8080/` to access the GraphQL Playground and interact with the API.

## Docker

Build the Docker image:
```bash
docker build -t gqlgen-app .
```

Run the container:
```bash
docker run -p 8080:8080 --env-file .env gqlgen-app
```

## CI/CD Pipeline

The project includes a GitHub Actions pipeline (`.github/workflows/deploy.yml`) that triggers on pushes to the `main` branch.

### Pipeline Steps:
1. **Build & Test**: Checks out code, sets up Go, downloads dependencies, and runs tests.
2. **Docker Build & Push**: Builds the Docker image and pushes it to Docker Hub.
3. **Deploy**: SSHes into the production server, pulls the latest image, and restarts the container.

### GitHub Secrets Required:
To use the CI/CD pipeline, add the following secrets to your GitHub repository:
- `DOCKER_USERNAME`
- `DOCKER_PASSWORD`
- `SERVER_HOST`
- `SERVER_USER`
- `SERVER_SSH_KEY`

## Deployment Options
You can deploy this application using various platforms:
- **Render/Railway**: Connect your GitHub repo, define the Build Command (`go build -o server server.go`), Start Command (`./server`), and set environment variables.
- **DigitalOcean/VPS**: The GitHub Action included automatically deploys via SSH to any Linux server running Docker.

## Testing with GraphQL Playground
Once running, you can test the following:
1. **Signup / Login** (Mutation)
2. **Create Post** (Mutation)
3. **Subscription** `subscription { postCreated { id title } }` (Open a new tab to see real-time updates when a post is created)
