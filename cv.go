package main

import (
        "net/http"
        "fmt"
)

var port = 3000

func main() {
        fmt.Printf("Hosting CV listening at port %d", port)
        http.Handle("/", http.FileServer(http.Dir("./public")))
        http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
        
        
}
