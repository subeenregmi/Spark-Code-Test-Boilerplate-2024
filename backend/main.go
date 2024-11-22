package main

import (
	"my-project/api"
	"my-project/database"
	"net/http"
)

func main() {
    db := database.CreateNewDB()
    todoapp := api.TodoHandler{DB: db}

    http.Handle("/", &todoapp)
    http.ListenAndServe(":8080", nil)
}

