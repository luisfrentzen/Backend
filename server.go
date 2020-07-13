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
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	//db := pg.Connect(&pg.Options{
	//	Addr:     "ec2-52-200-48-116.compute-1.amazonaws.com:5432",
	//	User:     "pgogvipbnkibet",
	//	Password: "f81c58a58203d8adb2589461aedb28b997f4aa980c2bec1754f10c8d179112b3",
	//	Database: "dbln7qhliqr8to",
	//})

		opt, err := pg.ParseURL("postgres://pgogvipbnkibet:f81c58a58203d8adb2589461aedb28b997f4aa980c2bec1754f10c8d179112b3@ec2-52-200-48-116.compute-1.amazonaws.com:5432/dbln7qhliqr8to?sslmode=require")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
