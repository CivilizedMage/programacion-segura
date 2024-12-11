package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hola Mundo")
}

func main() {
// Servir archivos estáticos desde el directorio "static"
fs := http.FileServer(http.Dir("./web"))
http.Handle("/", fs)

    http.HandleFunc("/hello", helloHandler)
    //localhost = 127.0.0.1
    fmt.Println("Servidor iniciado en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

