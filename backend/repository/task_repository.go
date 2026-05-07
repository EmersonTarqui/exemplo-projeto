package repository

import (
	"database/sql"
	"fmt"
	"todo-list-echo/backend/model"
)

// TaskRepository estrutura o acesso aos dados
type TaskRepository struct {
	connection *sql.DB
}

// NewTaskRepository cria uma nova instância do repositório
func NewTaskRepository(connection *sql.DB) TaskRepository {
	return TaskRepository{connection: connection}
}

// CreateTask insere uma tarefa e retorna o ID gerado
func (tr *TaskRepository) CreateTask(task model.Task) (int, error) {
	var id int

	// 1. Preparamos a Query
	// 1. Usamos "?" para garantir que o banco receba apenas os valores dos campos, evitando SQL Injection.
	query, err := tr.connection.Prepare("INSERT INTO tasks (user_id, title, done) VALUES (?, ?, ?)")

	// Se der erro ao preparar a query (ex: erro de escrita no SQL)
	if err != nil {
		fmt.Println("Erro ao preparar a query:", err)
		return 0, err
	}
	defer query.Close()

	// 2. Executamos a inserção passando os valores da struct
	result, err := query.Exec(task.UserID, task.Title, task.Done)

	// Se der erro ao inserir os dados no banco
	if err != nil {
		fmt.Println("Erro ao executar a inserção:", err)
		return 0, err
	}

	// 3. pegamos o ID gerado através do LastInsertId
	lastId, err := result.LastInsertId()

	// Se der erro ao tentar recuperar o ID que o banco criou
	if err != nil {
		fmt.Println("Erro ao recuperar o ID gerado:", err)
		return 0, err
	}

	id = int(lastId)

	// Se der sucesso, devolve o ID da tarefa criada e o erro vazio (nil)
	return id, nil
}
