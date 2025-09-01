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

## Setup

```bash
# Clone the repo
git clone https://github.com/VeluraDoc/Velura-Backend-Main.git
cd Velura-Backend-Main

# Install Go dependencies
go mod tidy

# (Optional) Setup Python env for PDF → DOCX conversion
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt

# Generate swagger docs
make doc

# Run with Docker
docker-compose up --build
```

## Local Run (without Docker)

```bash
make dev
```

This will:

1. Generate Swagger docs
2. Run the server on `http://localhost:8080`

## PDF → DOCX Conversion

This project uses a small Python script (`scripts/convert_pdf.py`) with [pdf2docx](https://pypi.org/project/pdf2docx/) for PDF → Word conversion.

- Make sure Python ≥3.9 is installed.
- Install dependencies inside a virtualenv:

```bash
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

- Test conversion manually:

```bash
python scripts/convert_pdf.py -i ./data/input.pdf -o ./data/output.docx
```

The Go code calls this script automatically using `exec.Command`.

## Requirements

- Go 1.22+
- Python 3.9+
- Docker & Docker Compose

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
swag init --parseDependency --parseInternal -g cmd/main.go -o docs
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

## 📫 Contact

For collaborations, partnerships, or inquiries:

- 🌐 Website: [Velura](https://velura-open-source-r36i.vercel.app)
- 💼 LinkedIn: [LinkedIn Page](https://www.linkedin.com/company/velura-open-source)
- 📧 Email: AliMoradi0Business@gmail.com

---

## 📜 License

[Apache](./LICENSE) © Velura Team
