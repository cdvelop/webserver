package gatitos

import (
	"fmt"
	"net/http"
)

func ManejadorGatitos(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de la URL
	query := r.URL.Query()
	// Obtener el valor del parámetro 'name_cat'
	name_cat := query.Get("name_cat")

	// obtener el apelido del gatito
	last_name := query.Get("last_name")

	// Si 'name_cat' no está presente, puedes asignar un valor por defecto
	if name_cat == "" {
		name_cat = "un gatito sin nombre"
	}

	if last_name == "" {
		last_name = "un gatito sin apellido"
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
		<h1>Nombre del gatito: %s</h1>
		<h1>Apellido: %s</h1>
	</body>
	</html>
	`, name_cat, last_name) // Incluir la variable nameCat en el HTML
}
