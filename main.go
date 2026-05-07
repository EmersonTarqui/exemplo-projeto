package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list-echo/backend/database" // Importa a conexão

	"github.com/gorilla/mux"
)

func main() {
	// Liga o banco de dados
	database.Connect()

	// Cria o roteador (quem cuida das URLs)
	router := mux.NewRouter()

	// Rota de teste: quando acessar http://localhost:8080/ping, verá "pong"
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	}).Methods("GET")

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
