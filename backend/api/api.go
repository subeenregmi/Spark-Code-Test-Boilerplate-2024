package api

import (
	"encoding/json"
	"io"
	"my-project/database"
	"my-project/service"
	"net/http"
)

type TodoHandler struct {
    DB  *database.Database
}

// Represents both Request and Response for POST (create to do).
type TodoRR struct {
    Title        string  `json:"title"`
    Description  string  `json:"description"`
}

type GetTodoListResponse []TodoRR

func HandleHTTPError(w http.ResponseWriter, code int) {
    // Common API errors

    var emsg string

    switch code {
        case http.StatusBadRequest:
            emsg = "Body cannot be read or not in correct JSON format."

        case http.StatusMethodNotAllowed:
            emsg = "Request method should be either GET or POST."

        default:
            emsg = "Undefined error."
    }

    http.Error(w, emsg, code)
}

func ValidMethod(method string) bool {
    return method == http.MethodGet || method == http.MethodPost
}

func (h *TodoHandler) GetTodoList(body []byte, w http.ResponseWriter) {
    // Gets all the to do's and send a JSON array of to do's

    listModels := service.GetAllTodos(h.DB)
    list := make([]TodoRR, 0, len(listModels))

    for _, elem := range listModels {
        list = append(list, TodoRR{elem.Title, elem.Description})
    }

    jsonBody, _ := json.Marshal(&list)
    w.Write(jsonBody)
}

func (h *TodoHandler) CreateTodo(body []byte, w http.ResponseWriter) {
    // Checks if the request is valid, and if so creates a to do
    // in the database.

    var t TodoRR
    err := json.Unmarshal(body, &t)

    // Request body is not a valid json object.
    if err != nil {
        HandleHTTPError(w, http.StatusBadRequest)
        return
    }

    // Checks if title is present
    if t.Title == "" {
        HandleHTTPError(w, http.StatusBadRequest)
        return
    }

    tModel := database.TodoModel{Title: t.Title, Description: t.Description}
    service.AddTodo(h.DB, tModel)

    w.Write(body)
}

func (h *TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

    // Handling the CORs preflight request 
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }
    // Validate method, should only be POST or GET
    if !ValidMethod(r.Method) {
        HandleHTTPError(w, http.StatusMethodNotAllowed)
        return
    }

    // Validate body can be read
    body, err := io.ReadAll(r.Body)
    if err != nil {
        HandleHTTPError(w, http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    switch r.Method {
        case http.MethodGet:
            h.GetTodoList(body, w)
            return

        case http.MethodPost:
            h.CreateTodo(body, w)
            return
    }
}
