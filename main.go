package main

import (
	"log"

	"github.com/JohanHR/twittor/bd"
	"github.com/JohanHR/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
