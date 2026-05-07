package controller

import (
	"encoding/json"
	"net/http"
	"todo-list-echo/backend/model"
	"todo-list-echo/backend/repository"
)

type TaskController struct {
	repo repository.TaskRepository
}

func NewTaskController(repo repository.TaskRepository) TaskController {
	return TaskController{repo: repo}
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	// Pegamos o texto JSON que veio do site e transformamos em objeto para manipulação do Go.
	json.NewDecoder(r.Body).Decode(&task)

	// Se o nome da tarefa estiver vazio, a gente para tudo e avisa o erro.
	if task.Title == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{Message: "O título da tarefa não pode estar vazio"})
		return
	}

	// Com o objeto pronto e verificado, mandamos para o Repository salvar no banco.
	id, err := tc.repo.CreateTask(task)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{Message: "Erro ao salvar no banco"})
		return
	}

	task.ID = id

	// Devolvemos a tarefa pronta em formato de texto (JSON).
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Chamamos o método do Repository para buscar a lista de tarefas no banco.
	tasks, err := tc.repo.GetTasks()

	// Caso ocorra erro na consulta ao banco, retornamos o status 500 e a mensagem de erro.
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{Message: "Erro ao buscar a lista de tarefas no banco"})
		return
	}

	// Se a busca for bem-sucedida, retornamos a lista de tarefas em formato JSON com status 200.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
