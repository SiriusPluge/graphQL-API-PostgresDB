package main

import (
	"graphQL-API-PostgresDB/domain"
	"graphQL-API-PostgresDB/graph"
	"graphQL-API-PostgresDB/graph/generated"
	"graphQL-API-PostgresDB/postgres"
	"graphQL-API-PostgresDB/scripts"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	DB := postgres.ConnectDB()
	defer DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	Repo := postgres.DB{DB: DB}
	d := domain.NewDomain(Repo)


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Domain: d}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", scripts.AuthorizationTokenContextMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
