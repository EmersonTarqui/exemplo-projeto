package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	user := "root"
	password := ""
	host := "localhost"
	port := "3306"
	dbname := "todo_db"

	// Monta a string de conexão (DSN) usando as variáveis
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	// 1. Tenta abrir a conexão
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("ERRO: Configuração do banco inválida! ", err)
	}

	// 2. Tenta dar um ping no banco para ver se ele está vivo e responde
	if err := db.Ping(); err != nil {
		log.Fatal("ERRO: Não consegui falar com o MySQL! Verifique se o XAMPP está ligado. ", err)
	}

	fmt.Println("Conectado ao MySQL com sucesso!")
	return db
}
