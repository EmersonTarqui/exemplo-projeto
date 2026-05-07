# To-Do List Echo 🚀
Aplicação de Lista de Tarefas desenvolvida para a disciplina de Desenvolvimento Seguro para Sistemas de Internet (SENAC 2026).

## 📂 Estrutura do Diretório
frontend/: Interface Single Page Application (HTML, CSS e JS Puro).

backend/: API RESTful em Go.

db/: Scripts de migração e schema do MySQL.

## 🛠️ Tecnologias Utilizadas
Linguagem: Go 1.26.3

Roteador: Gorilla Mux

Banco de Dados: MySQL

Containerização: Podman (A ser configurado futuramente).

## 🚀 Como Rodar Localmente
Certifique-se de ter o MySQL ativo.

Execute o script em db/migrations.sql.

No diretório raiz, execute:

```bash
go run main.go
```

## 🔐 Padrões de Contribuição
Commits: Seguir Conventional Commits em Português Brasileiro.

Assinatura: Obrigatório o uso de chave GPG (-S).

DCO: Inclusão do Signed-off-by (-s).
