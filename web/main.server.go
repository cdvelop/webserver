package main

import (
	"net/http"
	"webserver/web/gatitos"
)

func main() {

	// Servir archivos estáticos solo si la ejecución es donde esta la carpeta public
	fs := http.FileServer(http.Dir("./public"))
	// servir independiente de la ruta del proyecto
	// fs := http.FileServer(http.Dir("c:/Users/SU_NOMBRE/XXXX/webserver/web/public"))

	// ruta a servir publica
	http.Handle("/", fs)

	// ruta a api
	http.HandleFunc("/gatitos", gatitos.ManejadorGatitos)

	// mensaje consola
	print("Servidor iniciado en http://localhost:8080")
	// inicia servidor
	http.ListenAndServe(":8080", nil)
}
