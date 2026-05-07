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
	//  Usamos "?" para garantir que o banco receba apenas os valores dos campos, evitando SQL Injection.
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

func (tr *TaskRepository) GetTasks() ([]model.Task, error) {
	// Prepara a query para pegar tudo usando
	rows, err := tr.connection.Query("SELECT id, user_id, title, done, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task // Criamos a lista vazia

	// Percorre os resultados do banco linha por linha
	for rows.Next() {
		var task model.Task
		// O Scan "copia" os dados do banco para dentro da nossa struct task em cada campo
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Done, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		// Pega a tarefa que acabamos de preencher e adiciona na nossa lista "tasks"
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(task model.Task) error {
	// Preparamos a Query
	query, err := tr.connection.Prepare("UPDATE tasks SET done = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer query.Close()

	// Executamos a alteração
	result, err := query.Exec(task.Done, task.ID)
	if err != nil {
		return err
	}

	// VERIFICAÇÃO: Perguntamos ao banco quantas linhas foram alteradas
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Se o banco disser que 0 linhas foram alteradas, significa que o ID não existe
	if rowsAffected == 0 {
		return fmt.Errorf("tarefa com ID %d não encontrada", task.ID)
	}

	return nil
}
