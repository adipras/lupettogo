<p align="center">
  <img src="assets/images.png" alt="LupettoGo logo" />
</p>

<h1 align="center">LupettoGo 🐺</h1>
<p align="center"><i>With the little wolf, no project is too big.</i></p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#usage">Usage</a> •
  <a href="#examples">Examples</a>
</p>

---

# LupettoGo

🐺 **LupettoGo** is a powerful CLI tool that scaffolds production-ready Golang SaaS starter projects with clean architecture patterns. Generate complete applications with database integration, testing infrastructure, Docker support, and CRUD modules in seconds.

## ✨ Features

### 🏗️ **Clean Architecture Foundation**
- **Layered structure**: Handlers → Services → Repositories → Models
- **Dependency injection** with proper interface separation
- **SOLID principles** implementation out of the box
- **Scalable project organization** for enterprise applications

### 🗄️ **Database Integration**
- **PostgreSQL** and **MySQL** support with GORM
- **Auto-migrations** and connection management
- **Repository pattern** for data access layer
- **Environment-based** database configuration

### 🚀 **Development-Ready Setup**
- **HTTP server** with Gin framework
- **Middleware support**: CORS, logging, recovery
- **Configuration management** with Viper
- **Environment variables** with `.env` support
- **Structured logging** configuration

### 🧪 **Testing Infrastructure**
- **Unit test examples** with mocks
- **Test coverage** setup and reporting
- **Testify integration** for assertions and mocks
- **CI/CD ready** test structure

### 📦 **DevOps & Deployment**
- **Docker containerization** with multi-stage builds
- **Makefile** with common development tasks
- **Git configuration** with proper `.gitignore`
- **Production-ready** Dockerfile

### ⚡ **CRUD Module Generation**
- **Complete CRUD operations** for any entity
- **Auto-generated**: models, repositories, services, handlers
- **RESTful API** endpoints with proper HTTP methods
- **Validation and error handling** included

## 📦 Installation

### Option 1: Using `go install` (Recommended)
```bash
go install github.com/adipras/lupettogo@latest
```

### Option 2: Download Binary from Releases
```bash
# Linux/macOS
curl -fsSL https://raw.githubusercontent.com/adipras/lupettogo/main/install.sh | bash

# Or download manually from GitHub Releases
# https://github.com/adipras/lupettogo/releases
```

### Option 3: Build from Source
```bash
git clone https://github.com/adipras/lupettogo.git
cd lupettogo
go build -o lupettogo main.go
# Optional: Move to PATH
sudo mv lupettogo /usr/local/bin/
```

### Verify Installation
```bash
lupettogo version
lupettogo doctor  # Check your development environment
```

## 🚀 Quick Start

### 1. Create a New Project
```bash
# Basic project with PostgreSQL
lupettogo init my-saas-app

# Advanced project with custom configuration
lupettogo init my-api --db mysql --with-auth --with-docker
```

### 2. Setup and Run
```bash
cd my-saas-app
cp .env.example .env
# Edit .env with your database credentials
go mod tidy
go run main.go
```

### 3. Visit Your API
```bash
curl http://localhost:8080/health
# Returns: {"status":"ok","message":"🐺 LupettoGo API is running","version":"v1"}

curl http://localhost:8080/api/v1/example
# Returns: {"message":"Hello from LupettoGo! 🐺","status":"success","data":{...}}
```

## 📖 Usage

### Project Generation Options

```bash
lupettogo init [project-name] [flags]
```

**Flags:**
- `--db string`: Database driver (`postgres`, `mysql`) - default: `postgres`
- `--with-auth`: Include authentication scaffolding - default: `false`
- `--with-docker`: Include Docker configuration - default: `true`  
- `--with-tests`: Include testing infrastructure - default: `true`

### Module Generation

Generate complete CRUD modules within your project:

```bash
# Navigate to your project directory
cd my-saas-app

# Generate a user module
lupettogo generate module user
# Creates: user.go, user_repository.go, user_service.go, user_handler.go + tests
```

### Other Commands

```bash
lupettogo doctor    # Check development environment
lupettogo version   # Show version information
lupettogo --help    # Show all commands and options
```

## 💡 Examples

### Basic SaaS Project
```bash
lupettogo init blog-api
cd blog-api
lupettogo generate module post
lupettogo generate module user
```

### E-commerce Backend
```bash
lupettogo init ecommerce-api --db mysql --with-auth
cd ecommerce-api
lupettogo generate module product
lupettogo generate module order
lupettogo generate module customer
```

### Microservice with Testing
```bash
lupettogo init user-service --with-tests --with-docker
cd user-service
make test-coverage
make docker-build
```

## 🏗️ Generated Project Structure

```
my-saas-app/
├── 📄 main.go                    # Application entry point
├── 🔧 .env.example              # Environment template
├── 🐳 Dockerfile                # Container configuration
├── 📋 Makefile                  # Development commands
├── 📚 README.md                 # Project documentation
└── 📁 internal/
    ├── ⚙️  config/              # Configuration management
    ├── 🗄️  database/            # Database connection & migrations
    ├── 🎮 handlers/             # HTTP controllers with REST API
    ├── 🔀 middleware/           # HTTP middleware (CORS, auth, etc.)
    ├── 📊 models/               # Data models with GORM
    ├── 💾 repositories/         # Data access layer
    ├── 🧠 services/             # Business logic layer
    └── 🌐 server/               # HTTP server setup
```

## 🔧 Development Commands

The generated projects include a comprehensive Makefile:

```bash
make build          # Build the application
make run            # Run the application  
make test           # Run all tests
make test-coverage  # Run tests with coverage report
make lint           # Run code linting
make docker-build   # Build Docker image
make docker-run     # Run in Docker container
```

## 🌟 Why LupettoGo?

- **⚡ Fast Setup**: Go from idea to running API in under 2 minutes
- **🏢 Enterprise-Ready**: Clean architecture patterns used by top companies
- **🔧 Configurable**: Choose your database, features, and deployment options
- **📈 Scalable**: Built for growth with proper separation of concerns
- **🧪 Test-Driven**: Comprehensive testing setup from day one
- **📦 DevOps Ready**: Docker, Makefile, and CI/CD friendly

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <strong>🐺 With the little wolf, no project is too big.</strong>
</p>
