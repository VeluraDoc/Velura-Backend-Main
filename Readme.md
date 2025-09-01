# Velura Backend 🚀

A high-performance **Golang backend** designed with clean architecture principles, modular structure, and a focus on extensibility.  
The **core feature** of Velura is **PDF → Word conversion** 📝⚡, with secure authentication, user management, and modern DevOps support (Docker, Swagger, Makefile).

---

## ✨ Features

- 🔐 **Authentication & Authorization**

  - JWT-based auth with middleware
  - Role-based access control (RBAC)

- 👤 **User Management**

  - Sign up / Sign in
  - Secure password hashing (bcrypt)
  - Get current user profile (`/user/me`)

- 📑 **File Conversion**

  - Core business logic: **Convert PDF → Word**
  - Simple API endpoint for document transformation
  - Built for speed & reliability

- 🧩 **Architecture**

  - `internal/module/{auth,user}` split into `handler`, `usecase`, `repository`, `dto`, `model`
  - Repository pattern with singleton + interface
  - Clean separation of concerns

- ⚙️ **Developer Friendly**
  - Swagger UI (`/swagger/index.html`)
  - Dockerized (`docker-compose.yml`)
  - Auto migrations (GORM)
  - Makefile commands for dev workflow

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- `swag` CLI (for docs):

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

### Setup

```bash
# Clone the repo
git clone https://github.com/VeluraDoc/Velura-Backend-Main.git
cd Velura-Backend-Main

# Install dependencies
go mod tidy

# Generate swagger docs
swag init

# Run with Docker
docker-compose up --build
```

### Local Run

```bash
make dev
```

This will:

1. Generate Swagger docs
2. Run the server on `http://localhost:8080`

---

## 🔑 API Endpoints (Highlights)

| Method | Endpoint        | Description                | Auth |
| ------ | --------------- | -------------------------- | ---- |
| POST   | `/auth/sign-up` | Register a new user        | ❌   |
| POST   | `/auth/login`   | Login with credentials     | ❌   |
| GET    | `/user/me`      | Get current user profile   | ✅   |
| POST   | `/convert/pdf`  | Convert PDF → Word (.docx) | ✅   |

> Full API documentation available at:  
> 👉 **`/swagger/index.html`**

---

## 📖 Swagger Docs

Auto-generated with [swaggo/swag](https://github.com/swaggo/swag).

Run:

```bash
swag init
```

Then visit:

```bash
http://localhost:8080/swagger/index.html
```

---

## 🧪 Tests (Planned)

- Unit tests for usecases with repository mocks
- Integration tests for auth & user endpoints
- PDF → Word conversion validation

---

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **ORM**: GORM
- **Auth**: JWT + Middleware
- **Docs**: Swagger (swaggo)
- **Infra**: Docker & Docker Compose
- **Build**: Makefile automation

---

## 📌 Roadmap

- [ ] Add file upload & conversion queue
- [ ] Improve error responses (standard DTO)
- [ ] Add multi-language support
- [ ] CI/CD integration with GitHub Actions

---

## 🏆 Why Velura?

Because it's:

- 🔥 **Fast** — built with Go
- 🧼 **Clean** — layered architecture
- 🛡️ **Secure** — bcrypt + JWT + role middleware
- 🧩 **Modular** — easy to extend (swap repo: GORM → Mongo, etc.)
- ⚡ **Practical** — solves a real problem: PDF → Word conversion

---

## 📜 License

[Apache](./LICENSE) © Velura Team
