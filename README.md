# Go Greeting API Beginner Toolkit

This repository is the working codebase for a Moringa AI Capstone project:

**Prompt-Powered Kickstart: Building a Beginner's Toolkit for Go**

The project is a small HTTP API written in Go. It helps beginners learn how to:

- Create a Go module.
- Build routes with the standard `net/http` package.
- Return plain text and JSON responses.
- Test HTTP handlers using Go's built-in testing tools.

## Project Structure

```text
.
├── .gitignore
├── README.md
├── TOOLKIT.md
├── go.mod
├── main.go
└── main_test.go
```

## Requirements

- Go 1.22 or newer
- A terminal
- A code editor such as VS Code

Install Go from the official download page:

https://go.dev/doc/install

Confirm the installation:

```bash
go version
```

## Run the App

From the project folder:

```bash
go run .
```

Expected terminal output:

```text
Go Greeting API running at http://localhost:8080
```

Open another terminal and test the API:

```bash
curl http://localhost:8080/
```

Expected output:

```text
Welcome to the Go Greeting API. Try /greet?name=Amina or /health.
```

Try the greeting endpoint:

```bash
curl "http://localhost:8080/greet?name=Amina"
```

Expected output:

```json
{
  "name": "Amina",
  "message": "Hello, Amina! Welcome to Go.",
  "timestamp": "2026-05-02T12:00:00Z"
}
```

The timestamp will be different every time because the app generates it when the request is made.

## Test a POST Request

```bash
curl -X POST http://localhost:8080/greet \
  -H "Content-Type: application/json" \
  -d '{"name":"Grace Hopper"}'
```

Expected output:

```json
{
  "name": "Grace Hopper",
  "message": "Hello, Grace Hopper! Welcome to Go.",
  "timestamp": "2026-05-02T12:00:00Z"
}
```

## Run Tests

```bash
go test ./...
```

Expected output:

```text
ok  	moringa-go-greeting-api	0.XXXs
```

## Verified Locally

This project was verified on May 2, 2026 with:

```text
go version go1.26.2 darwin/arm64
ok  	moringa-go-greeting-api	0.575s
```

The local API smoke test returned:

```text
GET /       -> Welcome to the Go Greeting API. Try /greet?name=Amina or /health.
GET /health -> {"status":"ok","app":"go-greeting-api"}
GET /greet?name=Amina -> {"name":"Amina","message":"Hello, Amina! Welcome to Go.","timestamp":"2026-05-02T19:41:48Z"}
POST /greet -> {"name":"Grace Hopper","message":"Hello, Grace Hopper! Welcome to Go.","timestamp":"2026-05-02T19:41:45Z"}
```

## Optional: Change the Port

The app uses port `8080` by default. To run it on another port:

```bash
PORT=3000 go run .
```

Then visit:

```text
http://localhost:3000
```

## Capstone Documentation

The full toolkit document is in [TOOLKIT.md](/Users/user/Desktop/projects/project/TOOLKIT.md). It includes setup steps, the minimal working example, AI prompt journal, common issues, peer testing checklist, and references.
