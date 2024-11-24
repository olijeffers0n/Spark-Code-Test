package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos []Todo
var mutex sync.Mutex

func main() {
	todos = []Todo{}

	http.HandleFunc("/", ToDoListHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Handle error
		fmt.Println("Error starting server: ", err)
	}
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {

	case http.MethodGet:
		// Return the list of todos
		mutex.Lock()
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(todos)
		if err != nil {
			// Handle error
			fmt.Println("Error encoding response: ", err)
		}
		mutex.Unlock()

	case http.MethodPost:
		// Create a new todo
		var todo Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			// Handle error
			fmt.Println("Error decoding request: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if todo.Title == "" || todo.Description == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		mutex.Lock()
		todos = append(todos, todo)
		mutex.Unlock()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todo)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
