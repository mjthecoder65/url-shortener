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
  "shortCode": "abc123",
  "createdAt": "2021-09-01T12:00:00Z",
  "updatedAt": "2021-09-01T12:00:00Z"
}
```

### 2. Retrieve Original URL

**Request:**

```
GET /api/v1/shorten/abc123
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
PUT /api/v1/shorten/abc123
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
DELETE /api/v1/shorten/abc123
```

**Response:**

```
204 No Content
```

### 5. Get URL Statistics

**Request:**

```
GET /api/v1/shorten/abc123/stats
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

## Terraform Infrastructure Setup

This project uses Terraform to provision and manage the infrastructure required for deploying the URL Shortener API on Google Cloud Platform (GCP). The Terraform configuration automates the creation of networking resources, a Redis instance, an Artifact Registry repository, and a service account for CI/CD deployment via GitHub Actions.

### Terraform Resources

The Terraform configuration provisions the following resources:

| Resource Type                         | Name                              | Description                                                                                                           |
| ------------------------------------- | --------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `google_project_service`              | `redis_api`                       | Enables the Redis API for Memorystore.                                                                                |
| `google_project_service`              | `vpc_access_api`                  | Enables the VPC Access API for serverless VPC connectors.                                                             |
| `google_compute_network`              | `var.vpc_name`                    | Custom VPC network for the application (no auto-created subnets).                                                     |
| `google_compute_subnetwork`           | `asia-seoul-subnet`               | Subnetwork for general resources in the specified region.                                                             |
| `google_compute_subnetwork`           | `serverless-vpc-connector`        | Subnetwork for the serverless VPC connector.                                                                          |
| `google_compute_firewall`             | `allow-internal`                  | Firewall rule allowing all TCP/UDP/ICMP traffic within VPC subnets.                                                   |
| `google_compute_firewall`             | `allow-connector`                 | Firewall rule allowing Redis traffic (port 6379) from the VPC connector.                                              |
| `google_compute_firewall`             | `allow-health-checks`             | Firewall rule allowing health checks from Google's IP ranges.                                                         |
| `google_vpc_access_connector`         | `shorturlconnector`               | Serverless VPC connector for connecting Cloud Run to the VPC (2-3 instances).                                         |
| `google_redis_instance`               | `redis`                           | Managed Redis instance for caching (size configurable via `redis_size`).                                              |
| `google_artifact_registry_repository` | `var.artifact_registry_repo_name` | Docker repository in Artifact Registry for storing container images.                                                  |
| `google_service_account`              | `github-actions-deployer`         | Service account for GitHub Actions to deploy to Cloud Run.                                                            |
| `google_project_iam_member`           | Multiple roles                    | IAM roles for the service account: `run.admin`, `artifactregistry.writer`, `iam.serviceAccountUser`, `storage.admin`. |
| `google_service_account_key`          | `github_deployer_key`             | Key for the service account, used in GitHub Actions secrets.                                                          |

### Prerequisites for Terraform

- **Terraform**: Install Terraform (>=1.0) on your local machine.
- **Google Cloud SDK**: Install and authenticate `gcloud` with your GCP account.
- **GCP Project**: A Google Cloud project with billing enabled.
- **Service Account**: A GCP service account with permissions to manage resources (e.g., `Owner` or specific IAM roles).

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
