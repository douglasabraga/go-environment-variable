package main
import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	envTest := os.Getenv("GO_ENV")
	fmt.Printf("%s\n", envTest)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}