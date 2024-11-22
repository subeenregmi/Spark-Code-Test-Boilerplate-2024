import React, { useEffect, useState } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';

function TodoForm({todos, setTodos} : {todos: TodoType[], setTodos: Function}) {
    // 'todos' state is set as prop to this form element

    async function handleSubmit(event: React.SyntheticEvent) {
        event.preventDefault(); // Allow CORs and stop page reset on submit
        
        const target = event.target as typeof event.target & {
            title: { value: string };
            description: { value: string };
        };
        const title = target.title.value;
        const description = target.description.value;

        try {
            // Create to-do
            const response = await fetch("http://localhost:8080/", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    "title": title,
                    "description": description
                })
            });
            const data = await response.json();

            if (response.ok) {
               console.log(data);
               let list = todos.slice();
               list.push(data);
               setTodos(list); // Updates the to do list after creating a new list
            }

        } catch (error) {
            console.error("Error occurred when trying to create a new todo: ", error);
        }
    }

    return (
        <form onSubmit={handleSubmit}>
        <input placeholder="Title" name="title" autoFocus={true} />
        <input placeholder="Description" name="description" />
        <button>Add Todo</button>
      </form>
    );
}

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);

  // Initially fetch todo
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const todos = await fetch('http://localhost:8080/');
        if (todos.status !== 200) {
          console.log('Error fetching data');
          return;
        }

        setTodos(await todos.json());
      } catch (e) {
        console.log('Could not connect to server. Ensure it is running. ' + e);
      }
    }

    fetchTodos()
  }, []);

  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo, i) =>
          <Todo
            key={i}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
        <TodoForm todos={todos} setTodos={setTodos} />
   </div>
  );
}

export default App;
