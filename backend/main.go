package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
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
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Return the list of todos
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode([]string{})
		if err != nil {
			// Handle error
			fmt.Println("Error encoding response: ", err)
		}
	}
}
