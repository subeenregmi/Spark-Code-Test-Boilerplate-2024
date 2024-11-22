package database

import "sync"

type Database struct {
    TodoList  map[int]TodoModel
    mu        sync.RWMutex // Mutex ensures thread safety
    Count     int
}

type TodoModel struct {
    Title        string
    Description  string
}

func (db *Database) Add(t TodoModel) {
    db.mu.Lock()
    defer db.mu.Unlock()

    db.Count++
    db.TodoList[db.Count] = t
}

func (db *Database) GetAll() []TodoModel {
    db.mu.RLock() // Reads do not block other reads.
    defer db.mu.RUnlock()

    list := make([]TodoModel, 0, db.Count)
    for _, v := range db.TodoList {
        list = append(list, v)
    }

    return list
}

func CreateNewDB() *Database {
    db := Database {
        TodoList: make(map[int]TodoModel),
    }
    return &db
}
