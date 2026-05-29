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
<img width="252" height="195" alt="image" src="https://github.com/user-attachments/assets/b8867f17-a782-4ca5-a032-6280b61b7d33" />

## Build Executable

```bash
go build -o todo.exe
```

Run application:

```bash
todo.exe add "Buy milk"
```
```bash
todo.exe list
```
```bash
todo.exe modify 1 "Get car"
```
```bash
todo.exe done 1
```
```bash
todo.exe remove 1
```

## After using the application

### Stop Containers (Keep Data)

```bash
docker compose down
```

### Cleanup (Remove All Data)

```bash
docker compose down -v
```
