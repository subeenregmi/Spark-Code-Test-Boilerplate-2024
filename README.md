# To-do list application 
# Design Decisions
For the backend, I separated it into three packages, `api`, `service` and `database`. This made it easy to test and ensured that implementing futures changes to any part did not affect the rest. 

For the frontend, I separated the `<form>` tag into its own react component.

## Assumptions
The following where assumed to be true when completing this project:

- When creating a new to-do, `title` had to be a non-empty string.
- When an invalid request was received, the response was an string representation of the error.

## Model Representation
- To-do
	There are two structs that represent a single to-do, one struct `TodoRR` defined in the `api` package and `TodoModel` defined in the `database` package. Although not implemented, having a separate database representation allows you stores information that may not be necessary in the creation of a to-do and in the `GET` response (e.g: time of creation).

## Database
The struct `Database` defines our database, and has the following fields:
- `TodoList` - mapping from `int` to `TodoModel`
- `mu` - a `sync.RWMutex` ensuring that concurrent reading/writing to the database is safe
- `count` - stores number of to-do's.

Having an integer as a key for `TodoList` enables you to store  two different to-do's with the same title. Using a `sync.RWMutex` ensures that concurrent reads to the database do not block each other which can help the scalability of this solution. `count` stores the length of the list and by using the current length as a key we ensure that all to-dos will have a unique key.

In a future implementation, I would like to add caching to this database something which is not practical when the current database is just a map.

# Setup
If not already installed, please install the following:
1. Go ([install instructions](https://go.dev/doc/install))
2. Node ([download page](https://nodejs.org/en/download))

We have tested this with Node 20. You may have issues if you try to use a different version

# Running
Open two separate terminals - one for the React app and one for the golang API

## Golang API
1. In the first terminal, change to the backend directory (`cd backend`)
2. Run `go run main.go` to start the API server

This must be running for the frontend to work
When you make a change, you must stop the server (`ctrl-c` in the terminal), and restart it with `go run main.go`

## React App
1. In the second terminal, change to the frontend directory (`cd frontend`)
2. Run `npm start` to start the React app server
3. If it doesn't open automatically, open [http://localhost:3000](http://localhost:3000) to view your website

Leave this running. It will automatically update when you make any changes
