package main

import (
	"log"
	"net/http"
	"os"

    "database/sql"
    _ "github.com/mattn/go-sqlite3" // Use an anonymous import for the SQLite driver

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/monchipi/ex-myserver/graph/generated"
    "github.com/monchipi/ex-myserver/resolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

  //Initialize the SQLite database
  //Open("sqlite3","file:chatdatabase.db?cache=shared&mode=memory")
  db, err := sql.Open("sqlite3","./new.db")
  if err != nil{
    log.Fatalf("Failed to open the SQLite database: %v", err)
  }
  defer db.Close()

  srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
