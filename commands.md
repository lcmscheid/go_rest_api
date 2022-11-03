### Create *rest* db on postgres
- Log to postgres user: `sudo -iu postgres`
- Create the DB: `createdb rest`

### Execute the following command from the root working directory of the project
`migrate create -ext sql -dir migrations -seq create_todo_table`

### When using Migrate CLI we need to pass to database URL. Let's export it to a variable for convenience:
`export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/rest?sslmode=disable'`

### Then run migration files by executing the command:
`migrate -database ${POSTGRESQL_URL} -path migrations up`

### To rollback the previously executed migration run the following command:
`migrate -database ${POSTGRESQL_URL} -path migrations down`