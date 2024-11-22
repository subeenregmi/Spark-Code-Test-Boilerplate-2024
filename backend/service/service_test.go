package service 

import (
    "testing"
    "my-project/database"
)

func TestAddTodo(t *testing.T) {
    db := database.CreateNewDB()
    tm := database.TodoModel{Title: "Hello world!", Description: "Life is beautiful!"}

    AddTodo(db, tm)
    
    if db.Count != 1 {
        t.Errorf("AddTodo(db, tm) has failed. Expected: 1. Got: %d", db.Count)
    }

    if v := db.TodoList[1]; v != tm {
        t.Errorf("AddTodo(db, tm) has failed. Expected: %v. Got: %v", tm, v)
    }
}

func TestGetAllTodos(t *testing.T) {
    db := database.CreateNewDB()
    tm := database.TodoModel{Title: "Hello world!", Description: "Life is beautiful!"}

    for i := 0; i < 100; i++ {
        AddTodo(db, tm)
    }

    if v := GetAllTodos(db); len(v) != 100 {
        t.Errorf("GetAllTodos(db) has failed. Expected: 100. Got: %d", len(v))
    }
}
