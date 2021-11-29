package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v9"
	"github.com/phifertoo/go_graphql_server/graph"
	"github.com/phifertoo/go_graphql_server/graph/generated"
	"github.com/phifertoo/go_graphql_server/postgres"

	_ "github.com/99designs/gqlgen/cmd"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "lance",
		Password: "",
		Database: "meetmeup_dev4",
	})

	defer DB.Close()

	fmt.Println(DB)

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// need to pass the repos  into the server
	c := generated.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
