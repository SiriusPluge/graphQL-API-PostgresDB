# graphQL-API-PostgresDB сервис
Описание тестового задание можно посмотреть в файле "README-2.md"

# Для запуска сервера
- go run server/server.go

# Docker - postgresSQL

1) sudo docker run --name=postgres -e POSTGRES_PASSWORD='postgres' -p 7323:5432 -d postgres
2) migrate -path "postgres/db/migrations" -database "postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable" up
- migrate install: `go get -u github.com/golang-migrate/migrate`
4) sudo cat ./postgres/db/seeds/seeds.sql | psql postgres://postgres:postgres@localhost:7323


# API
## Взаимодествие с API:
- http://localhost:8080/ (in the browser)
- http://localhost:8080/query (in the GraphQL IDE)

### Примеры запросов

Получить список продуктов
```graphql
query {
  products {
    id
    name
  }
}
```

Отправить смс с кодом на указанный номер (для теста просто писать код в терминале). Запрос вернет null (все ок) или ошибку.
```graphql
mutation {
  requestSignInCode(input: { phone: "799999999" }) {
    message
  }
}
```

Авторизация с номер+код, результатом является токен или ошибка
```graphql
mutation {
  signInByCode(input: { phone: "799999999", code: "0000" }) {
    ... on SignInPayload {
      token
      viewer {
        user {
          phone
        }
      }
    }
    ... on ErrorPayload {
      message
    }
  }
}
```

С токеном можно получить данные пользователя. Токен передавать через куки или Authorization заголовок (лучше заголовок тк это поддерживает GraphqlIDE встроенная)
```graphql # любая команда прогона миграций
query {
  viewer {
    user {
      phone
    }
  }
}
```