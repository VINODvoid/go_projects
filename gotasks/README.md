# 🧩 Go Task Manager API

A simple **REST API built with Go** that manages tasks.  
This project is part of the learning path to understand **backend development** using Go — starting from CLI apps to HTTP servers.

---

## 🚀 Features (Phase 1 Complete)

✅ Implemented:
- `GET /tasks` → returns all tasks in JSON format  
- Basic Go web server using `net/http`  
- Error handling with idiomatic Go (`if err != nil`)  
- Organized project structure for scalability

🔜 Next (Coming Phases):
- `POST /tasks` → add a new task  
- `DELETE /tasks/:id` → remove a task  
- `PUT /tasks/:id` → update a task  
- Persistent storage using JSON or a database  

---

## ⚙️ Project Setup

1. **Initialize project:**
   ```bash
   mkdir go-task-manager
   cd go-task-manager
   go mod init go-task-manager
   ```

2. **Create `main.go`:**
   ```bash
   touch main.go
   ```

3. **Add this code:**
   ```go
   package main

   import (
       "encoding/json"
       "fmt"
       "net/http"
   )

   type Task struct {
       ID    int    `json:"id"`
       Title string `json:"title"`
       Done  bool   `json:"done"`
   }

   var tasks = []Task{
       {ID: 1, Title: "Learn Go basics", Done: true},
       {ID: 2, Title: "Build a REST API", Done: false},
   }

   func getTasks(w http.ResponseWriter, r *http.Request) {
       w.Header().Set("Content-Type", "application/json")

       err := json.NewEncoder(w).Encode(tasks)
       if err != nil {
           http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
           return
       }
   }

   func main() {
       http.HandleFunc("/tasks", getTasks)

       fmt.Println("Server running on http://localhost:8080")
       err := http.ListenAndServe(":8080", nil)
       if err != nil {
           fmt.Println("Server error:", err)
       }
   }
   ```

---

## ▶️ Run the Server

```bash
go run main.go
```

📍 You’ll see:
```
Server running on http://localhost:8080
```

---

## 🧪 Test the API

Open your browser or use `curl`:

```bash
curl http://localhost:8080/tasks
```

**Response:**
```json
[
  {"id":1,"title":"Learn Go basics","done":true},
  {"id":2,"title":"Build a REST API","done":false}
]
```

---

## 💡 Key Concepts Learned

### 1. `net/http` Basics
- `http.HandleFunc` → registers a route  
- `http.ResponseWriter` → writes data back to the client  
- `*http.Request` → reads request data (method, body, params)

### 2. Structs & JSON
- Define data models using Go `struct`s  
- Use `encoding/json` for serialization

### 3. Error Handling
```go
if err != nil {
    // handle it
}
```
- `err` is **not nil** → something went wrong  
- Always handle errors immediately and return to prevent crashes  

### 4. Pointers (`*http.Request`)
- The `*` means “pointer to” — Go passes the **reference** of the request, not a copy  
- You can safely use it without deep pointer knowledge for now

---

## 📘 Example Folder Structure
```
go-task-manager/
│
├── main.go
└── go.mod
```

---

## 🧠 Next Steps
1. Add `POST /tasks` → accept JSON body and append new task  
2. Add `DELETE /tasks/:id` → remove specific tasks  
3. Learn about routers (`gorilla/mux`)  
4. Integrate persistent storage (JSON file / PostgreSQL)

---

## 👨‍💻 Author
**Kalki**  
Learning Go by building real-world projects.  
Focus: Web Development • REST APIs • Scalable Backend Design

---

**Project Status:** 🟢 Phase 1 — `GET /tasks` completed  
**Next Goal:** 🔵 `POST /tasks` to add tasks dynamically
