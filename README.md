# Todo App

I have implemented the backend of the todo app in golang and the frontend in React. The backend is a simple REST API 
that allows you to create and read todos. The frontend is a simple todo list that allows you to interact with the API.

I have decided to stick to the OpenAPI specification provided in the task description. If I was given more freedom I would:

1) Extend the `To-Do` model to include more fields such as a unique ID, a due date, and a priority level
2) Using the unique ID, I would implement the ability delete Todos. Currently, you would only be able to 
delete a Todo by searching through the list of Todos and deleting the one you want to delete which is not efficient and doesn't handle duplicate data well. 

If I was to extend the project further I would:
1) Implement a database to store the Todos. Currently, the To-Dos are stored in memory and are lost when the server is
stopped. I would use a database such as PostgreSQL or SQLite to store the Todos.
2) Implement a user system. Currently, the Todos are shared between all users. I would implement a user system so that
each user has their own list of Todos.

## Setup
If not already installed, please install the following:
1. Go ([install instructions](https://go.dev/doc/install))
2. Node ([download page](https://nodejs.org/en/download))

We have tested this with Node 20. You may have issues if you try to use a different version

## Running
Open two separate terminals - one for the React app and one for the golang API

### Golang API
1. In the first terminal, change to the backend directory (`cd backend`)
2. Run `go run main.go` to start the API server

This must be running for the frontend to work
When you make a change, you must stop the server (`ctrl-c` in the terminal), and restart it with `go run main.go`

### React App
1. In the second terminal, change to the frontend directory (`cd frontend`)
2. Run `npm start` to start the React app server
3. If it doesn't open automatically, open [http://localhost:3000](http://localhost:3000) to view your website

Leave this running. It will automatically update when you make any changes
