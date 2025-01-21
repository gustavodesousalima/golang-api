# Golang API

API RESTful desenvolvida em Golang. A API permite criar, ler, atualizar e deletar usuários em um banco de dados MySQL.

- Aquele famoso CRUD 

## Funcionalidades

- **Criar Usuário**: Adiciona um novo usuário ao banco de dados.
- **Listar Usuários**: Retorna todos os usuários cadastrados.
- **Atualizar Usuário**: Atualiza as informações de um usuário existente.
- **Deletar Usuário**: Remove um usuário do banco de dados.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação utilizada para desenvolver a API.
- **MySQL**: Banco de dados relacional utilizado para armazenar os dados dos usuários.
- **Gorilla Mux**: Pacote utilizado para roteamento das requisições HTTP.
- **Go-SQL-Driver**: Driver utilizado para conectar ao banco de dados MySQL.
- **Godotenv**: Pacote utilizado para carregar variáveis de ambiente a partir de um arquivo `.env`.

## Configuração

1. Clone o repositório:

```sh
git clone https://github.com/gustavodesousalima/golang-api.git
cd golang-api
```

2. Crie um arquivo `.env` na pasta app com as seguintes variáveis:

```sh
DATABASE="SUA_URL_DO_BANCO_DE_DADOS"
API_KEY="SUA_API_KEY"
```

3. Instale as dependências:

```sh
go mod tidy
```

4. Entrar no pasta cmd/app e execute o servidor:

```sh
cd cmd/app

go run main.go
```

## Endpoints

### Criar Usuário
- POST /users

### Listar Usuários
- GET /users

### Atualizar Usuário
- PUT /users/{id}

### Deletar Usuário
- DELETE /users/{id}

---

## Exemplo de SQl 

```sh
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    createdat varchar(255) NOT NULL,
    active BOOLEAN DEFAULT TRUE,
);  
```

## Autenticação

- A API utiliza uma chave de API para autenticação. A chave deve ser enviada no cabeçalho Authorization de cada requisição.

### Exemplo de Request

```sh
curl -X GET "http://localhost:8080/users" \
-H "Authorization: SUA_API_KEY" \
-H "Content-Type: application/json" \
```