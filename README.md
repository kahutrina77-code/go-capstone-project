# Getting Started with Go (Golang) 🇰🇪

A simple REST API web server built with Go as part of 
Moringa School Capstone Project.

---

## Description
This project demonstrates a basic HTTP web server built 
using Go's standard library. It handles multiple routes 
and returns both plain text and JSON responses.

---

## Requirements
- Go 1.21.6+
- WSL (Ubuntu) or Linux/Mac terminal
- VS Code (recommended editor)

---

## How to Run

### 1. Clone the repository
```bash
git clone git@github.com:kahutrina77-code/go-capstone-project.git
```

### 2. Enter the project folder
```bash
cd go-capstone-project
```

### 3. Run the server
```bash
go run main.go
```

### 4. Visit in browser
```
http://localhost:8080
```

---

## API Routes

| Route | Description | Response Type |
|-------|-------------|---------------|
| / | Welcome message | Plain text |
| /api | JSON response | JSON |
| /about | Project information | Plain text |
| /time | Current server time | Plain text |

---

## Test with curl
```bash
# Test home route
curl http://localhost:8080

# Test API route (JSON)
curl http://localhost:8080/api

# Test about route
curl http://localhost:8080/about

# Test time route
curl http://localhost:8080/time
```

---

## Expected Output

### /api route returns:
```json
{
  "message": "Hello from Nairobi!",
  "location": "Nairobi, Kenya 🇰🇪",
  "status": "success",
  "timestamp": "2026-02-17 14:30:00"
}
```

---

## Project Structure
```
go-capstone/
├── main.go        # Main application code
├── go.mod         # Go module file
├── .gitignore     # Git ignore rules
├── README.md      # Project documentation
└── TOOLKIT.md     # Beginner's toolkit document
```

---

## Built With
- Go 1.21.6
- WSL Ubuntu on Windows
- VS Code
- AI Tools: Claude, GitHub Copilot, ai.moringaschool.com

---

## Author
**TRINA LUSENO** - Moringa School 2026

---

## License
MIT
