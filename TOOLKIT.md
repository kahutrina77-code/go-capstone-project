# Getting Started with Go (Golang) – A Beginner's Toolkit 🇰🇪

**By:** Trina Luseno
**School:** Moringa School
**Date:** February 2026

---

## 1. Title & Objective

**Technology Chosen:** Go (Golang)

**Why I Chose It:**
I chose Go because I wanted to learn something completely 
different from JavaScript (my primary language). Go is a 
compiled, statically typed language used by major companies 
like Google, Uber, and Netflix. It challenged me to think 
differently about programming and introduced me to backend 
development concepts.

**End Goal:**
Build a simple REST API web server in Go that:
- Handles multiple HTTP routes (/, /api, /about, /time)
- Returns JSON responses
- Runs on localhost:8080
- Can be tested via browser and terminal

---

## 2. Quick Summary of Go

**What is it?**
Go (also called Golang) is an open-source, compiled, 
statically typed programming language created by Google 
in 2009. It was designed for simplicity, speed, and 
reliability in large-scale software systems.

**Where is it used?**
- Backend web servers and REST APIs
- Cloud infrastructure (Docker and Kubernetes are built in Go)
- Microservices architecture
- Command-line tools
- DevOps tooling

**One Real-World Example:**
Uber uses Go to handle millions of requests per second 
for their ride-matching and pricing services. Go's speed 
and built-in concurrency make it perfect for high-traffic 
systems like Uber's backend.

---

## 3. System Requirements

- **OS:** Windows 11 with WSL (Ubuntu/Linux)
- **Editor:** VS Code
- **Go Version:** 1.21.6
- **Terminal:** WSL Terminal
- **Browser:** Any (Chrome, Firefox, Edge) for testing
- **No external packages needed**
  (project uses Go's built-in standard library only)
- **AI Tools Used:** Claude (claude.ai), GitHub Copilot

---

## 4. Installation & Setup Instructions

### Step 1: Update your system
```bash
sudo apt update && sudo apt upgrade -y
```

### Step 2: Download Go
```bash
wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
```

### Step 3: Install Go
```bash
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
```

### Step 4: Set up PATH
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Step 5: Verify Installation
```bash
go version
# Expected output: go version go1.21.6 linux/amd64
```

### Step 6: Create Project Folder
```bash
mkdir go-capstone
cd go-capstone
go mod init go-capstone
touch main.go
code .
```

---

## 5. Minimal Working Example

**What it does:**
A REST API web server with 4 routes:
- `/`      → Welcome message
- `/api`   → Returns JSON data
- `/about` → Project information
- `/time`  → Current server time

**Code:**
```go
package main

import (
	"encoding/json" // For JSON encoding
	"fmt"           // For printing to console/browser
	"net/http"      // For HTTP server functionality
	"time"          // For getting current time
)

// Response defines the structure of our JSON response
// similar to an object in JavaScript
type Response struct {
	Message   string `json:"message"`
	Location  string `json:"location"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// homeHandler handles the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go API! 🇰🇪\nVisit /api for JSON\nVisit /about for project info\nVisit /time for current time")
}

// apiHandler handles "/api" route and returns JSON
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Tell browser to expect JSON
	w.Header().Set("Content-Type", "application/json")

	// Create response object
	response := Response{
		Message:   "Hello from Nairobi!",
		Location:  "Nairobi, Kenya 🇰🇪",
		Status:    "success",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Convert to JSON and send
	json.NewEncoder(w).Encode(response)
}

// aboutHandler handles "/about" route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Capstone Project\nBuilt by: [Your Name]\nSchool: Moringa School\nTech: Go 1.21.6")
}

// timeHandler handles "/time" route
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Monday, 02 January 2006 - 15:04:05")
	fmt.Fprintf(w, "Current Time: %s", currentTime)
}

func main() {
	// Register all routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/time", timeHandler)

	// Start server on port 8080
	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

**How to run:**
```bash
go run main.go
```

**Expected terminal output:**
```
🚀 Server running on http://localhost:8080
```

**Expected browser output at /api:**
```json
{
  "message": "Hello from Nairobi!",
  "location": "Nairobi, Kenya 🇰🇪",
  "status": "success",
  "timestamp": "2026-02-17 14:30:00"
}
```

---

## 6. AI Prompt Journal

### Prompt 1: Understanding Go
**Platform:** Claude (claude.ai)
**Date:** 17/02/2026

**Prompt Used:**
"What is Go programming language? Explain what it is,
why developers use it, how it differs from JavaScript,
and why it's good for beginners"

**AI Response Summary:**
Go is a fast, simple programming language created by Google in 2009, designed for building servers, APIs, and cloud tools. Developers choose it because it's nearly as fast as C++, handles thousands of simultaneous tasks with ease, and keeps codebases clean and readable. Unlike JavaScript, Go is statically typed and compiled, meaning bugs are caught before your code ever runs rather than at runtime. For JavaScript developers, it's an ideal second language — minimal syntax to learn, excellent built-in tools, and habits that make you a sharper programmer overall.

**My Evaluation:**
 ⭐⭐⭐⭐⭐- I was able to understand what go is and able to evaluate if i can learn and understand it

**What I Did Next:**
Proceeded to install Go on WSL terminal

---

### Prompt 2: Installation on WSL
**Platform:** Claude (claude.ai)
**Date:** 17/02/2026

**Prompt Used:**
"I am using WSL on Windows and want to install Go.
Give me step by step installation instructions
including PATH setup and verification"

**AI Response Summary:**
1. **Update Ubuntu** — Run `sudo apt update && sudo apt upgrade -y`
2. **Download Go** — Run `wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz`
3. **Install Go** — Run `sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz`
4. **Set PATH** — Add these lines to `~/.bashrc` via `nano ~/.bashrc`, then run `source ~/.bashrc`
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=$HOME/go
   ```
5. **Verify** — Run `go version` and you should see `go version go1.23.4 linux/amd64`

**My Evaluation:**
⭐⭐⭐⭐⭐-Claude provided clear steps which made it easy for go setup

**What I Did Next:**
Successfully installed Go and verified with go version command

---

### Prompt 3: Hello World Explanation
**Platform:** Claude (claude.ai)
**Date:** 17/02/2026

**Prompt Used:**
"I just ran my first Go program with fmt.Println.
Explain what package main means, what import fmt does,
what func main() is, and how it differs from JavaScript"

**AI Response Summary:**
1. **`package main`** — Tells Go this file is a runnable program, not a reusable library. Every Go file must declare a package.
2. **`import "fmt"`** — Brings in Go's built-in formatting package for printing. Like `require()` in Node.js but with no installation needed.
3. **`func main()`** — The entry point Go automatically runs first, like an engine ignition. Unlike JavaScript, you never call it yourself.
4. **`fmt.Println`** — Prints text to the terminal with a new line, just like `console.log()` in JavaScript but requires the `fmt` package explicitly.
5. **JS difference** — JavaScript runs dynamically top-to-bottom with no required structure; Go enforces a strict entry point, explicit imports, and won't compile if you import something you don't use.

**My Evaluation:**
 ⭐⭐⭐⭐⭐- I was fully able to understand the function and what each command does.

**What I Did Next:**
Built on this knowledge to create a REST API

---

### Prompt 4: Building REST API
**Platform:** Claude (claude.ai)
**Date:** 17/02/2026

**Prompt Used:**
"I want to build a simple REST API in Go with routes
for /, /api, /about and /time. Show me complete code
with comments. I am a beginner from JavaScript background"

**AI Response Summary:**
• **Complete working code** — A single `main.go` file with 4 routes (`/`, `/api`, `/about`, `/time`) using only Go's standard library, ready to run with `go run main.go`.

• **Every line explained** — Detailed comments throughout comparing Go to JavaScript (e.g., `package main` vs no required structure in JS, `http.HandleFunc` vs `app.get()`, `json.Encoder` vs `JSON.stringify()`).

• **Structs and JSON** — Shows how to define typed data structures with `type Response struct` and use `json:"fieldname"` tags to control JSON output keys, plus a reusable `sendJSON()` helper function.


**My Evaluation:**
 ⭐⭐⭐⭐

**What I Did Next:**
Copied the code, tested all routes in browser and curl

---

### Prompt 5: Understanding the Code
**Platform:** Claude (claude.ai)
**Date:** 17/02/2026

**Prompt Used:**
"Explain my Go REST API code in detail. What is a struct?
What do the backtick tags mean? What is http.ResponseWriter?
How does JSON encoding work in Go?"

**AI Response Summary:**
• **Structs are typed data blueprints** — Like defining an object's shape before using it. `type HomeResponse struct { Success bool; Message string }` means "this data will always have these exact fields with these exact types" (unlike JavaScript where objects can have any shape).

• **Backtick tags control JSON output** — The `json:"success"` tags tell Go what key names to use when converting structs to JSON. Without them, Go would output capitalized field names like `"Success"` instead of `"success"`.

• **http.ResponseWriter is the response object** — It's Go's equivalent of `res` in Express.js. You write INTO it using methods like `w.Header().Set()` and `w.WriteHeader()` to build the HTTP response that gets sent back to the client.

• **JSON encoding converts structs to JSON strings** — `json.NewEncoder(w).Encode(myStruct)` reads your struct, converts it to JSON format using the backtick tags, and writes it directly into the response writer. It's like doing `JSON.stringify(obj)` and `res.send()` in one step.

• **Route handlers follow a strict signature** — Every handler function must be `func(w http.ResponseWriter, r *http.Request)` where `w` lets you write the response and `r` contains the incoming request data (URL, headers, body). Go won't compile if you use a different signature.

**My Evaluation:**
 ⭐⭐⭐⭐⭐- I was able to understand the code

**What I Did Next:**
Fully understood the code and proceeded to document it

---

## 7. Common Issues & Fixes

### Issue 1: PATH not set after installation
**Error:**
```
command not found: go
```
**Fix:**
```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```
**Why it happened:**
Go's binary location was not added to system PATH automatically

---

### Issue 2: Port already in use
**Error:**
```
listen tcp :8080: bind: address already in use
```
**Fix:**
```bash
# Find what's using port 8080
sudo lsof -i :8080

# Kill the process
kill -9 [PID number]
```
**Reference:** https://stackoverflow.com/questions/11583562

---

### Issue 3: VS Code not opening from WSL
**Error:**
```
code: command not found
```
**Fix:**
Install the WSL extension in VS Code then run:
```bash
code .
```
**Reference:** https://code.visualstudio.com/docs/remote/wsl

---

## 8. References

### Official Documentation
- Go Official Docs: https://go.dev/doc/
- Go Standard Library: https://pkg.go.dev/std
- Go Tour (Interactive): https://go.dev/tour/
- Go HTTP Package: https://pkg.go.dev/net/http

### AI Tools Used
- Claude AI: https://claude.ai
- GitHub Copilot: https://github.com/features/copilot

### WSL Resources
- Microsoft WSL Docs: https://learn.microsoft.com/en-us/windows/wsl/

### Helpful Resources
- Go by Example: https://gobyexample.com/
- Stack Overflow Go: https://stackoverflow.com/questions/tagged/go