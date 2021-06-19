package main

/*
import (
	"log"
	"net/http"
	"os"

	hackernews "github.com/quakenroll/go_gql_hackernews/graph/generated"
	"github.com/quakenroll/go_gql_hackernews/internal/pkg/db/migrations/mysql"

	//"github.com/quakenroll/go_gql_hackernews/internal/pkg/db/migrations/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	//"github.com/quakenroll/go_gql_hackernews/internal/pkg/db/migrations/mysql"
)
*/
import (

	//"github.com/quakenroll/go_gql_hackernews/internal/auth"
	//_ "github.com/quakenroll/go_gql_hackernews/internal/auth"
	"log"
	"net/http"
	"os"

	"github.com/quakenroll/go_gql_hackernews/graph"
	gqlgen_hackernews "github.com/quakenroll/go_gql_hackernews/graph/generated"
	"github.com/quakenroll/go_gql_hackernews/internal/auth"
	database "github.com/quakenroll/go_gql_hackernews/internal/pkg/db/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware()) // this sets the handler that should be processed before any normal handler's routines

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(gqlgen_hackernews.NewExecutableSchema(gqlgen_hackernews.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
