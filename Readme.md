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
```

### Local Run

```bash
# To run cli version
make start-cli

# To run api version
make start-api

# Or you can simply run the cli version using docker
# Your pdf file must be in /data dir
docker build -f Dockerfile.cli -t velura-cli .
docker run --rm -v "$PWD/data:/data" velura-cli -i /data/{file}.pdf -o /data/{file}.docx

# Example:
docker run --rm -v "$PWD/data:/data" velura-cli -i /data/input.pdf -o /data/output.docx
```

- By default, the CLI container generates a basic `.env` file inside `/app`
- You can mount your own `.env` using `-v "$PWD/.env:/app/.env"` if needed

- CLI runs as a one-shot container for direct conversions
- API runs persistently at: `http://localhost:8080`

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
