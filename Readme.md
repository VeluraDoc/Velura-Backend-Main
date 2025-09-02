# Velura Documents Service 📑⚡

A **modular document conversion service** built with Go.  
Designed to be **fast, extensible, and developer-friendly**, with a clean architecture that makes it easy to add new converters in the future.

The current core feature: **Convert PDF → Word (.docx)**.  
Planned: Word → PDF, CSV → PDF, PDF → CSV, and more.

---

## ✨ Features

- 📑 **File Conversion**

  - ✅ PDF → Word (.docx)
  - 🔜 Word → PDF
  - 🔜 CSV → PDF
  - 🔜 PDF → CSV

- 🧩 **Architecture**

  - Clean architecture (`handler`, `usecase`, `repository`, `dto`)
  - Separate conversion backends (LibreOffice, Python, custom parsers)
  - Easy to extend with new formats

- ⚙️ **Developer Friendly**
  - REST API + CLI support
  - Swagger UI (`/swagger/index.html`)
  - Dockerized (`docker-compose.yml`)
  - Makefile for dev workflow

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Python ≥3.9 (for `pdf2docx` backend, optional)
- `swag` CLI for Swagger:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Setup

```bash
# Clone the repo
git clone https://github.com/VeluraOpenSource/Velura_Documents_Service.git
cd Velura_Documents_Service

# Install Go dependencies
go mod tidy

# (Optional) Setup Python env for PDF → DOCX conversion
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt

# Generate Swagger docs
make doc

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

## 🧪 Tests (Planned)

- Unit tests for conversion usecases
- Integration tests for `/convert/*` endpoints
- Output validation (e.g., Word file opens correctly)

---

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **Conversion Backends**:
  - PDF → Word: Python `pdf2docx` (default)
  - Planned: LibreOffice headless, custom parsers
- **Docs**: Swagger (swaggo)
- **Infra**: Docker & Docker Compose
- **Build**: Makefile automation

---

## 📌 Roadmap

- [ ] Add Word → PDF conversion
- [ ] Add CSV ↔ PDF conversions
- [ ] Improve error handling & validation
- [ ] Queue system for large conversions
- [ ] Multi-language support (UTF-8/RTL docs)
- [ ] CI/CD with GitHub Actions

---

## 🏆 Why Velura Documents?

- 🚀 **Fast** — Go backend with minimal overhead
- 🧼 **Clean** — modular, layered architecture
- 🧩 **Extensible** — add new formats easily
- 🛡️ **Reliable** — production-ready tooling

---

## 📫 Contact

- 🌐 Website: [Velura](https://velura-open-source-r36i.vercel.app)
- 💼 LinkedIn: [Velura LinkedIn](https://www.linkedin.com/company/velura-open-source)
- 📧 Email: AliMoradi0Business@gmail.com

---

## 📜 License

[Apache-2.0](./LICENSE) © Velura Team
