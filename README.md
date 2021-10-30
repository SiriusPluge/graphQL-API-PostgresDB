# graphQL-API-PostgresDB


sudo docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 7323:5432 -d postgres



migrate -path ./schema -database 'postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable' up
- https://github.com/golang-migrate/migrate



sudo docker exec -it <?> /bin/bash

select * from users;

INSERT INTO users VALUES (3, '83333333333');
