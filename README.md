# apix 🚀

**apix** is a fast, developer-friendly CLI tool for testing HTTP APIs with built-in authentication, environment management, and configuration support — written in Go.

---

## ✨ Features

* ⚡ Fast and lightweight CLI for HTTP requests
* 🔐 Built-in authentication support (Bearer tokens, API keys, etc.)
* 🌍 Environment-based configuration (`.env` support)
* 🧩 Modular command structure
* 📦 Clean architecture using Go best practices
* 🛠 Supports common HTTP methods (GET, POST, DELETE)

---

## 📁 Project Structure

```
apix/
├── cmd/                # CLI commands
│   ├── auth.go
│   ├── config.go
│   ├── delete.go
│   ├── env.go
│   ├── get.go
│   ├── post.go
│   └── root.go
│
├── internal/           # Internal application logic
│   ├── client/         # HTTP client logic
│   │   └── http_client.go
│   └── config/         # Config management
│       └── config.go
│
├── main.go             # Entry point
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

---

## 📦 Installation

### Using Go

```bash
go install github.com/sahshad/apix@latest
```

Make sure `$GOPATH/bin` is in your PATH.

---

## 🚀 Usage

### GET request

```bash
apix get https://api.example.com/users
```

### POST request

```bash
apix post https://api.example.com/users -d '{"name":"John"}'
```

### DELETE request

```bash
apix delete https://api.example.com/users/1
```

## 📄 License

This project is licensed under the MIT License.
