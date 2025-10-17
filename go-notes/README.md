# 🗒️ Go Notes API

A full-stack backend API built in **Go (Golang)** with a **SQLite database**, supporting full **CRUD (Create, Read, Update, Delete)** operations on notes.  
This project is part of a practical Go learning track — combining **REST APIs**, **databases**, and **clean routing** using **Gorilla Mux**.

---

## 🚀 Features

✅ `GET /notes` → Fetch all notes  
✅ `POST /notes` → Add a new note  
✅ `PUT /notes/{id}` → Update a note by ID  
✅ `DELETE /notes/{id}` → Delete a note by ID  
✅ Persistent storage with **SQLite**  
✅ Clean routing via **Gorilla Mux**  
✅ Structured error handling  

---

## 🧠 Concepts Learned

- Go’s `net/http` server
- REST API design patterns
- JSON encoding/decoding (`encoding/json`)
- Database integration with `database/sql`
- Prepared statements and `Exec`
- URL parameters via `mux.Vars`
- Error handling & resource cleanup (`defer`)
- CRUD operations with SQL

---

## ⚙️ Setup

### 1️⃣ Clone & Initialize

```bash
git clone <repo_url>
cd go-notes
go mod init go-notes
```

### 2️⃣ Install Dependencies

```bash
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3
```

### 3️⃣ Run the Server

```bash
go run main.go
```

✅ Output:
```
✅ Database ready
🚀 Server running at http://localhost:8080
```

---

## 🗄️ Database

SQLite file: `notes.db`

**Table:**
| Column | Type | Description |
|---------|------|-------------|
| `id` | INTEGER | Auto-increment primary key |
| `title` | TEXT | Note title |
| `content` | TEXT | Note content |

---

## 🧩 API Reference

### 🟢 GET /notes
Returns all notes.

**Request**
```bash
curl http://localhost:8080/notes
```

**Response**
```json
[
  {"id":1,"title":"Learn Go","content":"Study database/sql"},
  {"id":2,"title":"Build APIs","content":"Use Gorilla Mux"}
]
```

---

### 🟡 POST /notes
Creates a new note.

**Request**
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"title":"New Note","content":"Go is awesome!"}' \
http://localhost:8080/notes
```

**Response**
```json
{"id":3,"title":"New Note","content":"Go is awesome!"}
```

---

### 🔵 PUT /notes/{id}
Updates an existing note.

**Request**
```bash
curl -X PUT -H "Content-Type: application/json" \
-d '{"title":"Updated Note","content":"Edited successfully"}' \
http://localhost:8080/notes/1
```

**Response**
```json
{"id":1,"title":"Updated Note","content":"Edited successfully"}
```

---

### 🔴 DELETE /notes/{id}
Deletes a note by ID.

**Request**
```bash
curl -X DELETE http://localhost:8080/notes/1
```

**Response**
```
204 No Content
```

---

## 📘 Folder Structure

```
go-notes/
│
├── main.go
├── go.mod
├── go.sum
└── notes.db
```

---

## 🧩 Key Learnings

| Concept | Description |
|----------|--------------|
| `gorilla/mux` | Clean and scalable routing |
| `mux.Vars(r)` | Extract path parameters |
| `database/sql` | Standard Go DB interface |
| `json.NewDecoder` | Parse JSON body from requests |
| `defer` | Guaranteed cleanup (closing DB, statements) |
| `strconv.Atoi` | Convert string ID → int |

---

## 🧱 Error Handling

```go
if err != nil {
    http.Error(w, "Something went wrong", http.StatusInternalServerError)
    return
}
```
Go’s idiomatic approach ensures safe, clear flow control — handle errors *immediately* and exit early.

---

## 📈 Example Workflow

```bash
# Add a note
curl -X POST -H "Content-Type: application/json" -d '{"title":"Golang","content":"Study structs"}' http://localhost:8080/notes

# Get all notes
curl http://localhost:8080/notes

# Update a note
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Advanced Go","content":"Learn mux"}' http://localhost:8080/notes/1

# Delete a note
curl -X DELETE http://localhost:8080/notes/1
```

---

## 🧭 Next Steps (Phase 3)

🚀 Improve architecture:
- Move DB, routes, and handlers into separate files  
- Add `.env` file for configuration  
- Add logging middleware  
- Optional: switch to PostgreSQL  

---

## 👨‍💻 Author
**Kalki**  
Learning Go by building real-world backend projects.  
Focus: Web Dev • APIs • Scalable Systems  

---

**Project Status:** 🟢 Completed (CRUD + SQLite)  
**Next Phase:** 🔵 Modular architecture + Environment setup
