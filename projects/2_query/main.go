package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", userHandler)
	fmt.Println("Server listening on port 9000")
	http.ListenAndServe(":9000", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		fmt.Fprintln(w, "Please provide a name using the 'name' parameter in the URL.")
		return
	}

	// Используйте переданное имя, а не "Salviya"
	greeting := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, greeting)
}
