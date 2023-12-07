package dependency_injection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func StartServer() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(myGreeterHandler)))
}
