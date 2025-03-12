# 🚀 URL Shortener API

A high-performance URL shortener built with **Golang**, **Gin**, **MongoDB**, and deployed on **Google Cloud Run**. Supports full CRUD operations and tracks URL statistics.

![Build Status](https://github.com/mjthecoder65/url-shortener/actions/workflows/ci.yml/badge.svg)

---

## ✨ Features

✅ Shorten long URLs in seconds  
✅ Retrieve the original URL from a shortcode  
✅ Update or delete shortened URLs  
✅ Track access statistics (clicks per URL)  
✅ Load tested for high performance  
✅ CI/CD pipeline with **GitHub Actions**  
✅ Cloud-native deployment with **Cloud Run**

---

## 🔧 Installation

### 1️⃣ Clone the repository

```sh
git clone https://github.com/mjthecoder10/url-shortener.git
cd url-shortener
```

### 2️⃣ Set up environment variables

Create a `.env` file with the following:

```plaintext
MONGO_URI=mongodb://localhost:27017
ALLOWED_CHARS=abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
SHORT_CODE_LENGTH=6
```

### 3️⃣ Run with Docker

```sh
docker-compose up --build
```

---

## 🚀 API Endpoints

### 1️⃣ Create a Short URL

```http
POST /shorten
```

**Request:**

```json
{
  "url": "https://www.example.com/long-url"
}
```

**Response:**

```json
{
  "id": "1",
  "shortCode": "abc123",
  "url": "https://www.example.com/long-url",
  "createdAt": "2025-03-12T12:00:00Z"
}
```

### 2️⃣ Retrieve Original URL

```http
GET /shorten/{shortCode}
```

**Response:**

```json
{
  "id": "1",
  "shortCode": "abc123",
  "url": "https://www.example.com/long-url",
  "accessCount": 10
}
```

### 3️⃣ Update a Short URL

```http
PUT /shorten/{shortCode}
```

**Request:**

```json
{
  "url": "https://www.example.com/updated-url"
}
```

### 4️⃣ Delete a Short URL

```http
DELETE /shorten/{shortCode}
```

### 5️⃣ Get URL Statistics

```http
GET /shorten/{shortCode}/stats
```

---

## 📊 Load Testing

Run a performance test using [k6](https://k6.io/):

```sh
k6 run load_test.js
```

---

## 🎯 Running Tests

```sh
go test ./...
```

---

## ☁️ Deployment

**To deploy on Cloud Run:**

```sh
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/url-shortener
gcloud run deploy url-shortener --image gcr.io/YOUR_PROJECT_ID/url-shortener --platform managed
```

---

## 📝 License

This project is licensed under the [MIT License](LICENSE).

---

## 👨‍💻 Author

**Michael Jordan Ngowi**

- GitHub: [@mjthecoder10](https://github.com/mjthecoder10)
- LinkedIn: [Michael Ngowi](https://www.linkedin.com/in/michael-ngowi/)
