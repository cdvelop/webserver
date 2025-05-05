package main

import (
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
	http.HandleFunc("/gatitos", gatitosHandler)

	// mensaje consola
	print("Servidor iniciado en http://localhost:8080")
	// inicia servidor
	http.ListenAndServe(":8080", nil)
}
