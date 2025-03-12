# URL Shortener API

A high-performance URL Shortener API built with Go, Gin, MongoDB, and deployed on Cloud Run. This API allows users to create, retrieve, update, delete, and track statistics of shortened URLs.

## Features

- Create short URLs with a unique code
- Retrieve original URLs from short codes
- Update existing short URLs
- Delete short URLs
- Track access statistics
- Dockerized for easy deployment
- CI/CD with GitHub Actions

## Tech Stack

- **Language**: Go
- **Framework**: Gin
- **Database**: MongoDB
- **Containerization**: Docker
- **CI/CD**: GitHub Actions
- **Deployment**: Google Cloud Run
- **Artifact Registry**: Google Artifact Registry

## Installation

### Prerequisites

- Go (>=1.18)
- Docker & Docker Compose
- MongoDB

### Clone Repository

```sh
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

### Run Locally

```sh
docker-compose up --build
```

### Run Tests

```sh
go test ./...
```

## API Documentation

### 1. Create Short URL

**Request:**

```
POST /shorten
Content-Type: application/json

{
  "url": "https://www.example.com/some/long/url"
}
```

**Response:**

```
201 Created
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 2. Retrieve Original URL

**Request:**

```
GET /shorten/abc123
```

**Response:**

```
200 OK
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 3. Update Short URL

**Request:**

```
PUT /shorten/abc123
Content-Type: application/json

{
  "url": "https://www.example.com/some/updated/url"
}
```

**Response:**

```
200 OK
{
  "id": "1",
  "url": "https://www.example.com/some/updated/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:30:00Z"
}
```

### 4. Delete Short URL

**Request:**

```
DELETE /shorten/abc123
```

**Response:**

```
204 No Content
```

### 5. Get URL Statistics

**Request:**

```
GET /shorten/abc123/stats
```

**Response:**

```
200 OK
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z",
  "accessCount": 10
}
```

## Deployment

### Build and Push Docker Image to Google Artifact Registry

```sh
gcloud auth configure-docker

docker build -t us-central1-docker.pkg.dev/your-project-id/url-shortener/url-shortener .
docker push us-central1-docker.pkg.dev/your-project-id/url-shortener/url-shortener
```

### Deploy to Cloud Run

```sh
gcloud run deploy url-shortener \
  --image us-central1-docker.pkg.dev/your-project-id/url-shortener/url-shortener \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

## License

This project is licensed under the MIT License.

## Author

[Your Name](https://github.com/your-username)
