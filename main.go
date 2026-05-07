package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list-echo/backend/controller"
	"todo-list-echo/backend/database"
	"todo-list-echo/backend/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Conexão com o Banco
	db := database.Connect()

	// Montagem das Camadas
	// o Controller recebe direto o Repository
	TaskRepository := repository.NewTaskRepository(db)
	TaskController := controller.NewTaskController(TaskRepository)

	// Configuração do Servidor/Roteador
	router := mux.NewRouter()

	// Rota de teste simples
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "pong"}`))
	}).Methods("GET")

	// Rotas principais
	router.HandleFunc("/tasks", TaskController.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", TaskController.GetTasks).Methods("GET")

	// 4. Start do Servidor
	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
