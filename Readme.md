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

Make sure you have the following installed:

- [Go 1.21+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [Python ≥3.9](https://www.python.org/) (only needed if using the `pdf2docx` backend locally)
- [GNU Make](https://www.gnu.org/software/make/) (for running `make` commands)
- `swag` CLI (for generating Swagger docs):

  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

---

## 🧪 Usage

### 🟢 API Server

```bash
# Clone and setup project
git clone https://github.com/VeluraOpenSource/Velura_Documents_Service.git
cd Velura_Documents_Service
go mod tidy

# Generate Swagger docs
make doc

# Start API server
make start-api

# API will be running on:
http://localhost:8080
```

---

### ⚙️ CLI Mode

```bash
# Build CLI using docker
docker build -f Dockerfile.cli -t velura-cli .

# Run CLI to convert a file
docker run --rm -v "$PWD/data:/data" velura-cli -i /data/input.pdf
```

- Your file must be in the `data/` directory.
- The CLI container automatically generates a basic `.env` file.
- You can mount your own `.env` using:

  ```bash
  -v "$PWD/.env:/app/.env"
  ```

---

### 📦 Using as a Go Package

To import and use the core functionality in your own Go project:

```bash
go get github.com/VeluraOpenSource/Velura_Documents_Service
```

Then in your code:

```go
import "github.com/VeluraOpenSource/Velura_Documents_Service/internal/usecase"

// Use conversion usecase here
```

> 🐍 **Note for PDF → DOCX:**
> The Go package uses a Python backend (via subprocess) for PDF → Word conversion.
> To enable it, you must:

1. Have **Python ≥3.9** installed
2. Install the required dependency:

```bash
pip install pdf2docx
```

3. Make sure `tools/pdf_convertor.py` is accessible via relative path

4. Set environment variables (or `.env` file) to define CLI behavior

---

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **Conversion Backends**:
  - PDF → Word: Python `pdf2docx` (default)
- **Docs**: Swagger (swaggo)
- **Infra**: Docker & Docker Compose
- **Build**: Makefile automation

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
