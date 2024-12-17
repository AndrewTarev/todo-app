package main

import todo "todo-app"

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
	}
}
