package main

import (
	"fmt"
	"net/http"
)

// Task model
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks []Task
var CURRENT_ID = 1

func TaskHandler(w http.ResponseWriter, r *http.Request){}
func getTasks(w http.ResponseWriter, r *http.Request) {}
func createTask(w http.ResponseWriter, r *http.Request){}
func updateTask(w http.ResponseWriter,r *http.Request){}
func deleteTask(w http.ResponseWriter, r *http.Request){}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to http server by Go")
	})
	http.HandleFunc("/tasks", getTasks)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}