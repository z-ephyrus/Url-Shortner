# URL Shortener (Go)

A simple and lightweight **URL Shortener service** built with **Golang**.  
It allows users to shorten long URLs into compact links and redirect them efficiently.  

This project is in its early stages — currently handling URL shortening and redirection in memory.  
Next steps include adding **database support** and **request validation**.

---

## 🚀 Features (Current)
- Shorten long URLs into short codes
- Redirect short URLs to their original destination
- In-memory storage (map) for shortened links
- Clean and modular Go code

---

## 📌 Roadmap
### ✅ Completed
- URL shortening logic
- Basic HTTP routing

### 🔜 Next Steps
- Add **database support** (PostgreSQL / SQLite)  
- Implement **URL validation** (using regex or libraries)  
- Add **expiration support** for links  
- Provide an **API with JSON responses**  
- Optional: **Custom short codes**

---

## ⚙️ Tech Stack
- **Language:** Go (Golang)  
- **Standard Library:** `net/http`, `encoding/json`  
- (Planned) **Database:** PostgreSQL or SQLite  
- (Planned) **Validation:** `go-playground/validator` or regex  

---

## 🛠️ Getting Started
### Prerequisites
- [Go](https://go.dev/dl/) 1.21+  

### Setup
```bash
# clone repo
git clone https://github.com/your-username/url-shortener.git
cd url-shortener

# install dependencies (if any)
go mod tidy

# run server
go run main.go
