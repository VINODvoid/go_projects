package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Task model
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks []Task
var CURRENT_ID = 1

func taskHandler(w http.ResponseWriter, r *http.Request){
	path := strings.Trim(r.URL.Path,"/")
	parts := strings.Split(path,"/")

	if len(parts) == 1{
		switch r.Method {
		case http.MethodGet:
			getTasks(w,r)
		case http.MethodPost:
			createTask(w,r)
		default:
			http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		}
		return
	}
	if len(parts) == 2{
		id, error := strconv.Atoi(parts[1])
		if error != nil{
			http.Error(w,"Invalid task ID",http.StatusBadRequest)
			return
		}
		switch r.Method{
		case http.MethodPut:
			updateTask(w,r,id)
		case http.MethodDelete:
			deleteTask(w,r,id)
		default:
			http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		}
		return
	}
	http.NotFound(w,r)
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasks)
}
func createTask(w http.ResponseWriter, r *http.Request){
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}
	newTask.ID = CURRENT_ID
	CURRENT_ID++
	if newTask.Status == ""{
		newTask.Status = "pending"
	}
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}
func updateTask(w http.ResponseWriter,r *http.Request, id int){
	var updateTask Task
	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil{
		http.Error(w,"Invalid Json",http.StatusBadRequest)
		return
	}

	for i , t := range tasks{
		if t.ID == id {
			if updateTask.Title != ""{
				tasks[i].Title = updateTask.Title
			}
			if updateTask.Status != ""{
				tasks[i].Status = updateTask.Status
			}
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.Error(w,"Task not found",http.StatusNotFound)

}
func deleteTask(w http.ResponseWriter, r *http.Request,id int){
	for i,t := range tasks{
		if t.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return 
		}
	}
	http.Error(w,"Task not found",http.StatusNotFound)
}


func main() {

	http.HandleFunc("/tasks", taskHandler)
	http.HandleFunc("/tasks/",taskHandler )
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}