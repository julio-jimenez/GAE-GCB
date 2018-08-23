package main
import (
        "log"
        "net/http"
)
func handler(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello, world!\n"))
}
func main() {
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}