package main

import (
	"Backend/graph"
	"Backend/graph/generated"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "postgres://pgogvipbnkibet:f81c58a58203d8adb2589461aedb28b997f4aa980c2bec1754f10c8d179112b3@ec2-52-200-48-116.compute-1.amazonaws.com:5432/dbln7qhliqr8to",
		User:     "pgogvipbnkibet",
		Password: "f81c58a58203d8adb2589461aedb28b997f4aa980c2bec1754f10c8d179112b3",
		Database: "dbln7qhliqr8to",
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
