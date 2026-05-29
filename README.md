# Go CLI Todo App

A simple command-line todo application built with Go and PostgreSQL.

## Features
- Add todo
- Modify todo
- List todos
- Delete todo
- Mark todo as done

## Tech Stack
- Go
- PostgreSQL
- Docker

## Installation

```bash
git clone https://github.com/yourname/go-cli-todo.git
cd go-cli-todo
```

### 1. Create .env
config .env.example to .env then create user/password
```bash
DB_HOST=localhost
DB_PORT=5432

POSTGRES_DB=todo_db
POSTGRES_USER=nameStrongUser
POSTGRES_PASSWORD=nameStrongPassword

PGADMIN_EMAIL=admin@test.com
PGADMIN_PASSWORD=admin
```
### 2. Start database

```bash
docker compose up -d
```
#### 3. Check docker run
```bash
docker ps
```
status should show like this
```
STATUS
Up 5 minutes
Up 5 minutes
```

### 4. Run application

```bash
go run . list
```

## Commands

### Add todo

```bash
go run . add "Buy milk"
```

### List todos

```bash
go run . list
```

### Modify todo

```bash
go run . modify 1 "do homework"
```

### Check todo

```bash
go run . done 1
```

### Delete todo

```bash
go run . remove 1
```

## Example Output
