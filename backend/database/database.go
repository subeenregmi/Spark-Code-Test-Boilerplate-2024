package database

import "sync"

type Database struct {
    TodoList  map[int]TodoModel
    mu        sync.RWMutex // Mutex ensures thread safety
    count     int
}

type TodoModel struct {
    Title        string
    Description  string
}

func (db *Database) Add(t TodoModel) {
    db.mu.Lock()
    defer db.mu.Unlock()

    db.count++
    db.TodoList[db.count] = t
}

func (db *Database) GetAll() []TodoModel {
    db.mu.RLock() // Reads do not block other reads.
    defer db.mu.RUnlock()

    list := make([]TodoModel, 0, db.count)
    for _, v := range db.TodoList {
        list = append(list, v)
    }

    return list
}
