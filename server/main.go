package main

import (
	"fmt"
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)

	var port string = ":8080"
	fmt.Println("Server corriendo en el puerto", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}
