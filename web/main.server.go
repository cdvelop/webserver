package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Servir archivos estáticos solo si la ejecución es donde esta la carpeta public
	// fs := http.FileServer(http.Dir("./public"))
	// servir independiente de la ruta del proyecto
	fs := http.FileServer(http.Dir("c:/Users/Cesar/Classes/webserver/web/public"))

	// ruta a servir publica
	http.Handle("/", fs)

	// ruta a api
	http.HandleFunc("/api1", Api)

	// mensaje consola
	print("Servidor iniciado en http://localhost:8080")
	// inicia servidor
	http.ListenAndServe(":8080", nil)
}

func Api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Pagina AIEP</title>
		<meta charset="UTF-8">
	</head>
	<body>
		<img src="/static/pino.jpg" alt="Logo AIEP" width="100" />
		<h1>API</h1>
	</body>
	</html>
	`)
}
