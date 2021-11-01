# graphQL-API-PostgresDB


sudo docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 7323:5432 -d postgres

- https://github.com/golang-migrate/migrate
migrate -path "postgres/migrations" -database "postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable" up


INSERT INTO users VALUES (3, '83333333333');
