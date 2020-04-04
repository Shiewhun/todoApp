package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

//Todo Struct (Model)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

//Initialize todos var as a slice of Todo struct
var todos []Todo

//Create Todo list item
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = rand.Intn(100) // Create Mock ID
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func main() {
	//Initialize Router
	r := mux.NewRouter()

	//Mock Data
	todos = append(todos, Todo{ID: 1, Task: "Coding", Completed: true})
	todos = append(todos, Todo{ID: 2, Task: "Read other books", Completed: false})

	//route handler/endpoints
	r.HandleFunc("/api/todo", getTodos).Methods("GET")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
