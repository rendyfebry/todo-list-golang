# TODO List CLI Golang

Simple Todo List app

## Prerequisite

- Golang 1.9
- Dep

## Install depedencies

```
dep ensure
```

## Install Package Globally

```
go get -u github.com/rendyfebry/todo-list-golang/...
```

## Usage

```
todos-go [command]
```

```
add         Add task
complete    Complete a task
delete      Delete task
help        Help about any command
list        List tasks
sync        Sync data
```

PS: If you didn't install globally, you need to change the command from `todos-go` to `go run bin/todos-go/main.go`

```
go run bin/todos-go/main.go list
```

## Docker

Docker Build

```
docker build -t todos-go .
```

Run in Docker

```
docker run --rm -ti todos-go list
```
