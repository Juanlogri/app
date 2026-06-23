package main

import (
	"log"

	"github.com/Juanlogri/app/bd"
	"github.com/Juanlogri/app/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
