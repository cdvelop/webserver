package main

import (
	"fmt"
	"net/http"
)

func gatitosHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de la URL
	query := r.URL.Query()
	// Obtener el valor del parámetro 'name_cat'
	nameCat := query.Get("name_cat")

	// Si 'name_cat' no está presente, puedes asignar un valor por defecto
	if nameCat == "" {
		nameCat = "un gatito sin nombre"
	}

	// Usar el valor en la respuesta
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Pagina AIEP</title>
		<meta charset="UTF-8">
	</head>
	<body>
		<img src="img/gatito.jpg" alt="Gatito" width="200" />
		<h1>API - Nombre del gatito: %s</h1>
	</body>
	</html>
	`, nameCat) // Incluir la variable nameCat en el HTML
}
