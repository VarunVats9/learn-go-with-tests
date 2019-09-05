package injection

import (
	"fmt"
	"io"
	"net/http"
)

// Greet takes writer and name, and prints out message to stdout
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler takes ResponseWriter and Request, and calls Greet func
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
