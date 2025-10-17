package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("✅ Database ready")
}

// GET /notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content FROM notes")
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// POST /notes
func addNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if note.Title == "" || note.Content == "" {
		http.Error(w, "Missing title or content", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO notes(title, content) VALUES(?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare insert", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(note.Title, note.Content)
	if err != nil {
		http.Error(w, "Failed to insert note", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	note.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DELETE /notes/{id}
func deleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("DELETE FROM notes WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare delete", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PUT /notes/{id}
func updateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note Note
	err = json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("UPDATE notes SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare update", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(note.Title, note.Content, id)
	if err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	note.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func main() {
	initDB()
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/notes", getNotes).Methods("GET")
	r.HandleFunc("/notes", addNote).Methods("POST")
	r.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/notes/{id}", updateNote).Methods("PUT")

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
