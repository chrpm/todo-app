# todo-app
API for a TODO application written in golang

## Using this application

### Endpoints
Path: `/tasks`
Method: `GET`
Optional Parameter: `completed=true|false` for filtering 
Desc: Get all tasks

Path: `/tasks/{id}`
Method: `GET`
Desc: Get task by id

Path: `/tasks`
Method: `POST`
Desc: Create a new todo task
Content-Type: application/json
Body: {"goal": "Get oil changed"}

Path: `/tasks/{id}`
Method: `PUT`
Desc: Update task by id
Content-Type: application/json
Body: {"goal": "Clean room", "completed": true|false} // completed optional defaults to false

Path: `/tasks/{id}`
Method: `DELETE`
Desc: Delete task at id


## Running this application

1. Set Required env vars
    - TODO_DBUSER="user"
    - TODO_DBPASSWORD="password"
    - TODO_DATABASE="/todo_app"
    - TODO_PORT=":80"
2. Create MySQL table using command in `db/task_table.sql` 
3. Run using `make run` 