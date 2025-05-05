package main

import (
	"net/http"
)

func main() {

	// Servir archivos estáticos desde el directorio "static" de la maquina
	fs := http.FileServer(http.Dir("./public"))

	// ruta a servir publica
	http.Handle("/", fs)

	// mensaje consola
	print("Servidor iniciado en http://localhost:8080")
	// inicia servidor
	http.ListenAndServe(":8080", nil)
}
