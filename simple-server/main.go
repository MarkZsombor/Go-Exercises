package main

// Runs an HTTP server that will show the phrase "Hello <name>" when you go to localhost:8000/<name>
import (
	"net/http"
	"strings"
)

// Get the URL param then forms a string to be displayed on page
func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + ". Thanks for trying my program!"

	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
