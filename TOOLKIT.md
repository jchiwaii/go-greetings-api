# Prompt-Powered Kickstart: Building a Beginner's Toolkit for Go

## 1. Title and Objective

**Chosen technology:** Go, also called Golang.

**Project title:** Go Greeting API Beginner Toolkit.

**Why I chose Go:** Go is a beginner-friendly compiled language with simple syntax, fast execution, and built-in tools for formatting, testing, and running code. It is not Python, Java, or JavaScript, so it fits the capstone requirement.

**End goal:** Build and run a minimal HTTP API that returns greeting messages as JSON. A beginner should be able to install Go, run the app, test two API endpoints, and understand the basic structure.

## 2. Quick Summary of Go

Go is an open-source programming language designed for building reliable and efficient software. It is commonly used for backend services, command-line tools, cloud systems, DevOps tools, and APIs.

Go programs are organized into modules and packages. A typical beginner project starts with a `go.mod` file and a `main.go` file. The `main` package is the entry point for a runnable Go application.

**Real-world example:** Many backend teams use Go to build APIs because Go includes a strong standard library. In this project, the API is built with Go's built-in `net/http` package, so no external web framework is needed.

## 3. System Requirements

| Requirement | Details |
| --- | --- |
| Operating system | macOS, Linux, or Windows |
| Editor | VS Code, GoLand, Vim, or any text editor |
| Language toolchain | Go 1.22 or newer |
| Terminal tools | `go`, `curl` |
| External packages | None |

Install Go from the official documentation:

https://go.dev/doc/install

Confirm that Go is installed:

```bash
go version
```

Expected output:

```text
go version go1.xx.x darwin/arm64
```

The exact version and operating system may be different depending on the computer.

## 4. Installation and Setup Instructions

### Step 1: Create or open the project folder

```bash
mkdir go-greeting-api
cd go-greeting-api
```

This repository already contains the files, so if you are using this codebase, just open the folder in your terminal.

### Step 2: Initialize a Go module

```bash
go mod init moringa-go-greeting-api
```

This creates a `go.mod` file. This project already has this file:

```go
module moringa-go-greeting-api

go 1.22
```

### Step 3: Create the API file

Create `main.go`. This file contains the web server and routes.

### Step 4: Run the app

```bash
go run .
```

Expected output:

```text
Go Greeting API running at http://localhost:8080
```

### Step 5: Test in another terminal

```bash
curl http://localhost:8080/
```

Expected output:

```text
Welcome to the Go Greeting API. Try /greet?name=Amina or /health.
```

Test the JSON greeting route:

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

The timestamp changes on every request.

### Step 6: Run automated tests

```bash
go test ./...
```

Expected output:

```text
ok  	moringa-go-greeting-api	0.XXXs
```

## 5. Minimal Working Example

The example is a small API with these endpoints:

| Endpoint | Method | Purpose |
| --- | --- | --- |
| `/` | GET | Returns a welcome message |
| `/health` | GET | Returns a JSON health check |
| `/greet?name=Amina` | GET | Returns a greeting using a query parameter |
| `/greet` | POST | Returns a greeting using a JSON request body |

Important part of `main.go`:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// greetingResponse describes the JSON response sent back to the user.
type greetingResponse struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// newRouter connects URL paths to handler functions.
func newRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/greet", greetHandler)
	return mux
}

// greetHandler supports GET /greet?name=Amina.
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	writeJSON(w, http.StatusOK, newGreeting(name))
}

// newGreeting cleans the name and creates the response message.
func newGreeting(name string) greetingResponse {
	cleanName := strings.Join(strings.Fields(name), " ")
	if cleanName == "" {
		cleanName = "World"
	}

	return greetingResponse{
		Name:      cleanName,
		Message:   fmt.Sprintf("Hello, %s! Welcome to Go.", cleanName),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}

// writeJSON sends a Go value as a JSON HTTP response.
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
```

The project `main.go` expands this minimal example by also supporting `POST /greet`, invalid JSON handling, method checks, and a security header.

## 6. AI Prompt Journal

Documentation reference: This section records the GenAI prompts used for learning, planning, implementation, debugging, and reflection. The prompts should be run or adapted in `ai.moringaschool.com` as required by the capstone instructions.

| Prompt used | Curriculum or reference link | AI response summary | Evaluation and reflection |
| --- | --- | --- | --- |
| "Act as a beginner-friendly Go tutor. Explain what Go is, where it is used, and why it is good for building small APIs." | Moringa AI portal: https://ai.moringaschool.com/ and official Go learning page: https://go.dev/learn | The AI explained Go as a compiled language used for backend services, APIs, cloud tools, and CLIs. | Helpful because it gave me a simple explanation I could reuse in the toolkit summary. I still verified the learning path with official Go docs. |
| "Give me a step-by-step guide to initialize a Go module and create a Hello World API using only the standard library." | Official Go getting started tutorial: https://go.dev/doc/tutorial/getting-started | The AI suggested using `go mod init`, `main.go`, `net/http`, and `go run .`. | Very helpful for scaffolding. I refined the result by adding `/health`, `/greet`, and JSON responses. |
| "Explain `net/http` in Go using beginner terms and show how `http.HandleFunc` connects a route to a function." | Go `net/http` package docs: https://pkg.go.dev/net/http | The AI explained that a handler receives a request and writes a response. | Helpful because it clarified the mental model of request, handler, and response. I used this to document the routes clearly. |
| "I am getting `zsh: command not found: go`. What does it mean and how do I fix it on macOS, Linux, or Windows?" | Official Go install guide: https://go.dev/doc/install | The AI explained that Go is not installed or not on the PATH. It recommended installing Go and reopening the terminal. | Helpful for the common issues section. The fix must always point back to the official install guide. |
| "Review this beginner Go API README for clarity. What steps are missing for someone running it for the first time?" | Capstone brief and README checklist | The AI recommended adding requirements, run commands, expected outputs, endpoint examples, and tests. | Helpful because it improved the documentation for a beginner audience. I added expected outputs and common fixes. |
| "Create a peer testing checklist for a beginner who has never run a Go API before." | Capstone testing and iteration requirement | The AI suggested checking installation, `go run .`, curl commands, error messages, and README clarity. | Helpful because it turned testing into a repeatable process rather than a vague instruction. |

### Prompt usage feedback

Using GenAI helped me move faster because I could ask targeted questions instead of searching randomly. The most useful prompts were specific and included the desired output, such as "use only the standard library" or "explain for a beginner." The AI was less useful when prompts were too broad, so I improved the prompts by asking for step-by-step setup, expected errors, and verification commands.

## 7. Common Issues and Fixes

| Issue | Cause | Fix |
| --- | --- | --- |
| `zsh: command not found: go` | Go is not installed or the terminal cannot find it. | Install Go from https://go.dev/doc/install, then close and reopen the terminal. Run `go version` again. |
| `go: cannot find main module` | The project does not have a `go.mod` file or the command is being run from the wrong folder. | Run commands from the project root. If starting from scratch, run `go mod init moringa-go-greeting-api`. |
| `listen tcp :8080: bind: address already in use` | Another app is already using port 8080. | Run on another port: `PORT=3000 go run .`, then use `http://localhost:3000`. |
| Browser shows `404 page not found` | The URL path does not match an existing route. | Use `/`, `/health`, or `/greet?name=Amina`. |
| POST request returns `invalid JSON body` | The JSON body is malformed. | Send valid JSON: `{"name":"Amina"}` and include `Content-Type: application/json`. |
| `curl: command not found` | Curl is missing or unavailable in the terminal. | Use a browser for GET routes, or install curl. On Windows, PowerShell can use `Invoke-WebRequest`. |

## 8. Testing and Iteration

### Automated test checklist

Run:

```bash
go test ./...
```

The tests check that:

- The home route returns a welcome message.
- `GET /greet?name=Amina` returns the correct JSON greeting.
- `GET /greet` defaults to `World`.
- `POST /greet` accepts a JSON body.
- Invalid JSON returns a `400 Bad Request`.
- Unsupported methods return `405 Method Not Allowed`.

### Local verification output

This project was verified on May 2, 2026.

```bash
go version
```

Output:

```text
go version go1.26.2 darwin/arm64
```

```bash
go test ./...
```

Output:

```text
ok  	moringa-go-greeting-api	0.575s
```

API smoke test results:

```text
GET /       -> Welcome to the Go Greeting API. Try /greet?name=Amina or /health.
GET /health -> {"status":"ok","app":"go-greeting-api"}
GET /greet?name=Amina -> {"name":"Amina","message":"Hello, Amina! Welcome to Go.","timestamp":"2026-05-02T19:41:48Z"}
POST /greet -> {"name":"Grace Hopper","message":"Hello, Grace Hopper! Welcome to Go.","timestamp":"2026-05-02T19:41:45Z"}
```

### Peer testing checklist

Ask a peer to follow only the README, without extra explanation.

| Task | Peer result | Notes |
| --- | --- | --- |
| Install or confirm Go with `go version` | Pass / Needs help | |
| Run `go run .` | Pass / Needs help | |
| Open `http://localhost:8080/` | Pass / Needs help | |
| Run `curl "http://localhost:8080/greet?name=Amina"` | Pass / Needs help | |
| Run the POST example | Pass / Needs help | |
| Run `go test ./...` | Pass / Needs help | |
| README was clear enough to follow | Yes / No | |

### Iteration notes

After peer testing, update this section with any feedback. Example improvements might include clearer install steps, an extra screenshot, or a shorter explanation of `go.mod`.

## 9. References

- Official Go install guide: https://go.dev/doc/install
- Official Go getting started tutorial: https://go.dev/doc/tutorial/getting-started
- Official Go learning page: https://go.dev/learn
- Go `net/http` package documentation: https://pkg.go.dev/net/http
- Go `encoding/json` package documentation: https://pkg.go.dev/encoding/json
- Effective Go: https://go.dev/doc/effective_go
- Moringa AI portal for prompts: https://ai.moringaschool.com/
