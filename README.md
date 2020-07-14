# Start PostgreSQL with Docker for using in Web-app
```$ docker run -d --name chitchat-db -e POSTGRES_PASSWORD=12345678 -e POSTGRES_USER=chitchat -e POSTGRES_DB=chitchat -e PGDATA=/var/lib/postgresql/data/pgdata -p 5432:5432 postgres```

# Build the web-app
```$ go build -o ./src/main.go```
