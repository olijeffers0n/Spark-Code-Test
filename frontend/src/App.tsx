import React, { useEffect, useState } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';

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

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const formData = new FormData(event.currentTarget);
        const newTodo = {
            title: formData.get('title') as string,
            description: formData.get('description') as string,
        };

        event.currentTarget.reset();

        try {
            const response = await fetch('http://localhost:8080', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(newTodo),
            });

            if (response.status === 200) {
                setTodos([...todos, newTodo]);
            } else {
                console.log('Error posting data');
            }
        } catch (e) {
            console.log('Could not connect to server. Ensure it is running. ' + e);
        }
    };

  return (
    <div className="app">
      <header className="app-header">
        <h1>Todo Application</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form onSubmit={handleSubmit}>
        <input placeholder="Title" name="title" autoFocus={true} />
        <input placeholder="Description" name="description" />
        <button>Add Todo</button>
      </form>
    </div>
  );
}

export default App;
