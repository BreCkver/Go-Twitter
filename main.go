package main

import (
	"log"

	"github.com/BreCkver/Go-Twitter/bd"
	"github.com/BreCkver/Go-Twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
