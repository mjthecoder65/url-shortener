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
git clone https://github.com/mjthecoder65/url-shortener.git
cd url-shortener
```

### Run Locally

```sh
docker-compose up -d --build --remove-orphans
```

### Run Tests

```sh
  make test
```

## API Documentation

# URL Shortener API

## API Endpoints

| Method | Endpoint                           | Description                      |
| ------ | ---------------------------------- | -------------------------------- |
| POST   | `/api/v1/shorten`                  | Create a new short URL           |
| GET    | `/api/v1/shorten/:shortCode`       | Retrieve original URL & redirect |
| PUT    | `/api/v1/shorten/:shortCode`       | Update an existing short URL     |
| DELETE | `/api/v1/shorten/:shortCode`       | Delete a short URL               |
| GET    | `/api/v1/shorten/:shortCode/stats` | Get access count statistics      |

### 1. Create Short URL

**Request:**

```
POST /api/v1/shorten
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
  "shortCode": "rt8WMa",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 2. Retrieve Original URL

**Request:**

```
GET /api/v1/shorten/rt8WMa
```

**Response:**

```
200 OK
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "rt8WMa",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 3. Update Short URL

**Request:**

```
PUT /api/v1/shorten/rt8WMa
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
  "shortCode": "rt8WMa",
  "createdAt": "2025-03-01T12:00:00Z",
  "updatedAt": "2021-03-12T12:30:00Z"
}
```

### 4. Delete Short URL

**Request:**

```
DELETE /api/v1/shorten/rt8WMa
```

**Response:**

```
204 No Content
```

### 5. Get URL Statistics

**Request:**

```
GET /api/v1/shorten/rt8WMa/stats
```

**Response:**

```
200 OK
{
  "id": "1",
  "url": "https://www.example.com/some/long/url",
  "shortCode": "rt8WMa",
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

[Michael Jordan](https://github.com/mjthecoder65)
