# graphQL-API-PostgresDB сервис

#Для запуска сервера
- go run server/server.go

#Docker - postgresSQL

1) sudo docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 7323:5432 -d postgres
2) migrate -path "postgres/db/migrations" -database "postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable" up
- migrate install: `go get -u github.com/golang-migrate/migrate`
4) sudo cat ./postgres/db/seeds/seeds.sql | psql postgres://postgres:postgres@localhost:7323


#API
## Взаимодествие с API:
- http://localhost:8080/ (in the browser)
- http://localhost:8080/query (in the GraphQL IDE)


