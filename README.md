## apix 🚀

**apix** is a modern, developer-focused CLI tool for testing and interacting with HTTP APIs directly from the terminal.

Built with Go, apix provides a fast and lightweight experience with support for authentication, environment-based configuration, multipart uploads, structured response rendering, and common HTTP workflows.

---

### ✨ Features

* ⚡ Fast and lightweight HTTP API client
* 🔐 Built-in authentication support
* 🌍 Environment-aware configuration management
* 📦 Support for common HTTP methods
  * GET
  * POST
  * PUT
  * PATCH
  * DELETE
* 📁 JSON payload and file-based request support
* 📤 Multipart form-data file uploads
* 📊 Structured response rendering with timing information
* 🧩 Modular and maintainable architecture
* 🛠 Built using idiomatic Go practices

---


### 📦 Installation

#### Using Go

```bash
go install github.com/sahshad/apix@latest
```

Make sure `$GOPATH/bin` is in your PATH.

---

### 🚀 Usage

#### GET request

```bash
apix get https://api.example.com/users
```

---

#### POST request

Inline JSON Payload
```bash
apix post /users -d '{"name":"John"}'
```

Request Body from File
```bash
apix post /users -f body.json
```

Multipart File Upload
```bash
apix post /upload -F file=document.pdf
```

---

#### PUT Request

Inline JSON Payload
```bash
apix put /users/1 -d '{"name":"Updated"}'
```

Request Body from File
```bash
apix put /users/1 -f body.json
```

Multipart PUT Upload
```bash
apix put /upload/1 -F file=document.pdf
```

---

#### PATCH Request

Inline JSON Payload
```bash
apix patch /users/1 -d '{"status":"active"}'
```

Request Body from File
```bash
apix patch /users/1 -f body.json
```

Multipart PATCH Upload
```bash
apix patch /upload/1 -F file=document.pdf
```

---

#### DELETE request

```bash
apix delete https://api.example.com/users/1
```
---


#### 🔐 Authentication


Manage bearer token authentication for API requests.

Example:

```bash
apix auth login
```

---

#### 🌍 Environment Management

Manage multiple API environments easily.

Example:

```bash
apix env set development=https://dev-api.example.com
```

---

#### ⚙️ Configuration

Manage CLI configuration values.

Example:

```bash
apix config set base_url=https://api.example.com
```

---

### 📁 Project Structure

```
sahshad-apix/
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── main.go
│
├── cmd/
│   ├── auth.go
│   ├── config.go
│   ├── delete.go
│   ├── env.go
│   ├── get.go
│   ├── patch.go
│   ├── post.go
│   ├── put.go
│   └── root.go
│
└── internal/
    ├── cli/
    │   ├── client_helper.go
    │   ├── file.go
    │   ├── format.go
    │   ├── header.go
    │   ├── response.go
    │   └── ui.go
    │
    ├── client/
    │   └── http_client.go
    │
    ├── config/
    │   └── config.go
    │
    └── types/
        └── http.go
```
---

### 🧱 Architecture

apix follows a modular internal architecture designed for maintainability and scalability.

-- `cmd/`

  CLI command definitions and argument handling

-- `internal/client`

   HTTP request execution and multipart handling

-- `internal/cli`

   Response formatting, rendering, and terminal utilities

-- `internal/config`

   Configuration and environment management

-- `internal/types`

   Shared application types

---

### 📄 License

This project is licensed under the MIT License.
