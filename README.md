# 🚧 URL Shortening Service - Under Construction

[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)](https://golang.org)
[![MongoDB](https://img.shields.io/badge/MongoDB-v2.3.0-green.svg)](https://www.mongodb.com)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> A simple and efficient URL shortening service built with Go and MongoDB. This project implements the [roadmap.sh URL Shortening Service](https://roadmap.sh/projects/url-shortening-service) specification.

## 🚀 Features

- ✅ Create short URLs from long URLs
- ✅ Retrieve original URLs using short codes
- ✅ Update existing short URLs
- ✅ Delete short URLs
- ✅ Track URL access statistics
- ✅ MongoDB persistence with indexing
- ✅ RESTful API design

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/shorten` | Create a new short URL |
| `GET` | `/shorten/{shortcode}` | Retrieve original URL and increment access count |
| `PUT` | `/shorten/{shortcode}` | Update an existing short URL |
| `DELETE` | `/shorten/{shortcode}` | Delete a short URL |
| `GET` | `/shorten/{shortcode}/stats` | Get URL statistics |

### Example Usage

**Create Short URL:**
```bash
curl -X POST http://localhost:8090/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.example.com/very/long/url"}'
```

**Get Original URL:**
```bash
curl http://localhost:8090/shorten/abcdefghij
```

**Update URL:**
```bash
curl -X PUT http://localhost:8090/shorten/abcdefghij \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.updated-example.com"}'
```

## 🛠️ Setup & Installation

### Prerequisites
- Go 1.24.2 or higher
- MongoDB instance running
- Environment variables configured

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/OJPARKINSON/short-url.git
   cd short-url/go
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up environment variables:**
   Create a `.env` file:
   ```env
   MONGODB_URI=mongodb://localhost:27017
   DB_NAME=shorturl
   COLLECTION_NAME=urls
   ```

4. **Run the application:**
   ```bash
   go run main.go
   ```

5. **Access the API:**
   - API: `http://localhost:8090`

## 📊 Data Structure

```json
{
  "url": "https://www.example.com",
  "shortCode": "abcdefghij",
  "createdAt": "2025-09-07T22:00:00Z",
  "updatedAt": "2025-09-07T22:00:00Z",
  "accessCount": 5
}
```

## 🚧 TODO & Known Issues

### Missing/Issues from roadmap.sh evaluation:

#### Response Format Issues:
- [ ] Add proper JSON structure with `id` field
- [ ] Fix HTTP status codes consistency
- [ ] Standardize error response format

#### Technical Issues:
- [ ] Fix `SingleResult.Acknowledged` issue in GetShorten function
- [ ] Implement URL validation for malformed URLs
- [ ] Add short code collision handling
- [ ] Improve error handling consistency

#### Enhancements:
- [ ] Add input validation middleware
- [ ] Implement rate limiting
- [ ] Add logging middleware
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add Docker support
- [ ] Add environment-specific configs
- [ ] Add health check endpoint

## 🛠️ Tech Stack

- **Language:** Go 1.24.2
- **Database:** MongoDB v2.3.0
- **HTTP Router:** Go's native `http.ServeMux`
- **Environment:** dotenv

### Dependencies

```go
require (
    go.mongodb.org/mongo-driver/v2 v2.3.0
)
```

## 📝 Project Structure

```
.
├── main.go              # Application entry point
├── handlers/
│   ├── shortenHandler.go # URL CRUD operations
│   └── statsHandler.go   # Statistics handler
├── db/
│   ├── connect.go       # Database connection
│   └── init.go          # Database initialization
├── go.mod              # Go module file
└── README.md           # This file
```

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Project specification from [roadmap.sh](https://roadmap.sh/projects/url-shortening-service)
- MongoDB Go Driver documentation

---

**Status:** 🚧 Under Active Development (85% Complete)