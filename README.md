# Alumni Portal Backend (Go)

This is the backend service for the **Alumni Portal Project**, built with **Go**.  
It provides APIs for authentication, gallery management, alumni information, and payments.  
Currently, the project is in its early stage with routes, modular structure, and setup ready.  
Upcoming features include **database support** (PostgreSQL) and **data validation**.

---

## 🚀 Features (Current)
- Go module setup with clean project structure
- Authentication routes (JWT-based in progress)
- Gallery system with event image metadata
- Admin dashboard APIs (scaffolded)
- Payment system (Stripe integration planned)

---

## 📌 Roadmap
### ✅ Completed
- Initial Go setup  
- API routing structure  
- Modular project layout  

### 🔜 Next Steps
- Add **database support** (PostgreSQL + BUN ORM)  
- Implement **request validation** (using libraries like `go-playground/validator`)  
- Role-based authentication (Admin / Mentor / User)  
- Newsletter system with subscriber management  
- File uploads for gallery/events  

---

## ⚙️ Tech Stack
- **Language:** Go (Golang)  
- **Frameworks/Libs (planned):**
  - [Bun](https://bun.uptrace.dev/) – PostgreSQL ORM
  - [go-playground/validator](https://github.com/go-playground/validator) – Validation
  - [JWT](https://github.com/golang-jwt/jwt) – Authentication
- **Database (planned):** PostgreSQL  
- **Payments (planned):** Stripe  

---

## 🛠️ Getting Started
### Prerequisites
- [Go](https://go.dev/dl/) 1.21+  
- (Future) PostgreSQL running locally or via Docker  

### Setup
```bash
# clone repo
git clone https://github.com/your-username/alumni-portal-backend.git
cd alumni-portal-backend

# install dependencies
go mod tidy

# run server
go run main.go
