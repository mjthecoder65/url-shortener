# URL Shortener API

A URL Shortener API built with Go, Gin, MongoDB, and Redis. This API allows users to create, retrieve, update, delete, and track statistics of shortened URLs. Short URLs are essential in today’s digital landscape as they simplify sharing, improve readability, and save space in character-limited platforms like social media. They also enable better tracking and analytics, making it easier to monitor user engagement. For high availability and fault tolerance, I deployed the service on Google Cloud Run.

This project was implemented based on the roadmap requirements [roadmap.sh](https://roadmap.sh/projects/url-shortening-service).

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
- **Caching**: Redis
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

### Setting Up Environment Variables

#### **For Local Development**

Create a `.env` file in the root directory and populate it with the required variables:

```sh
touch .env
```

**Example `.env` File:**

```env
APP_ENV=dev # prod, staging
MONGODB_URI=mongodb://localhost:27017/main
MONGODB_PASSWORD=your_mongodb_password
SERVER_ADDRESS=:8080
ALLOWED_CHARS=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
SHORT_CODE_LENGTH=6
```

Ensure MongoDB is running locally on port `27017` before starting the application.

### **Run Locally**

```sh
docker-compose up -d --build --remove-orphans
```

### **Run Tests**

```sh
make test
```

## API Documentation

| Method | Endpoint                           | Description                      |
| ------ | ---------------------------------- | -------------------------------- |
| POST   | `/api/v1/shorten`                  | Create a new short URL           |
| GET    | `/api/v1/shorten/:shortCode`       | Retrieve original URL & redirect |
| PUT    | `/api/v1/shorten/:shortCode`       | Update an existing short URL     |
| DELETE | `/api/v1/shorten/:shortCode`       | Delete a short URL               |
| GET    | `/api/v1/shorten/:shortCode/stats` | Get access count statistics      |

## **Setting Up Secrets and Variables in GitHub Actions**

To ensure your CI/CD pipeline runs smoothly, you need to set up **secrets** and **variables** in GitHub Actions.

### **1. Setting Up Secrets**

Go to your GitHub repository → **Settings** → **Secrets and variables** → **Actions** → **Secrets** → **New repository secret**

Add the following secrets:

- **`MONGODB_PASSWORD`** → Your MongoDB password
- **`MONGODB_URI`** → The connection string for MongoDB
- **`GCP_CREDENTIALS`** → JSON credentials for your Google Cloud service account

### **2. Setting Up Variables**

Go to **Settings** → **Secrets and variables** → **Actions** → **Variables** → **New repository variable**

Add the following variables:

- **`APP_ENV`** → `production` (or `development`)
- **`SERVER_ADDRESS`** → `:8080`
- **`ALLOWED_CHARS`** → `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
- **`SHORT_CODE_LENGTH`** → `6`
- **`ARTIFACT_REGISTRY_REGION`** → Your Google Cloud region (e.g., `us-central1`)
- **`GCP_PROJECT_ID`** → Your Google Cloud project ID
- **`ARTIFACT_REGISTRY_REPO`** → Name of your Google Artifact Registry repo
- **`CLOUD_RUN_SERVICE_ACCOUNT`** → Email of the service account with Cloud Run deployment permissions

## Deployment

### Build and Push Docker Image to Google Artifact Registry

```sh
# Set the following environment variables in your terminal. Replace the placeholders
# with your own values
export REGION=asia-northeast3
export ARTIFACT_REPOSITORY_NAME=url-shortener
export GOOGLE_PROJECT_ID=rock-elevator-453623-f5
export IMAGE_TAG=v1
export IMAGE_NAME="${REGION}-docker.pkg.dev/${GOOGLE_PROJECT_ID}/${ARTIFACT_REPOSITORY_NAME}/url-shortener:$IMAGE_TAG"

# Create Artifact repository
gcloud artifacts repositories create "$ARTIFACT_REPOSITORY_NAME" \
  --repository-format=docker \
  --location=$REGION

# Authenticate Docker to push images to Artifact Registry
gcloud auth configure-docker

# Build the Docker Image
docker build -t "$IMAGE_NAME" .

# Push Docker Image to Artifact Registry
docker push "$IMAGE_NAME"
```

### Deploy to Cloud Run

```sh
gcloud run deploy url-shortener \
  --image="$IMAGE_NAME" \
  --platform=managed \
  --region="$REGION" \
  --allow-unauthenticated
```

## License

This project is licensed under the MIT License.

## Author

[Michael Jordan](https://github.com/mjthecoder65)
