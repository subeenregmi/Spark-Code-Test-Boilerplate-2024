package service

import "my-project/database"

/*
    This file is pretty empty, and may seem redudant but in 
    the future if, for example if a word-filter was to be needed,
    it can be implemeted here.
*/

func AddTodo(db *database.Database, t database.TodoModel) {
    db.Add(t)
}

func GetAllTodos(db *database.Database) []database.TodoModel {
    list := db.GetAll()
    return list
}
