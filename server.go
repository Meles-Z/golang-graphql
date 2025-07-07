package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/meles-z/golang-graphql/app/generated"
	"github.com/meles-z/golang-graphql/app/infrastructure/db"
	"github.com/meles-z/golang-graphql/app/infrastructure/persistence"
	"github.com/meles-z/golang-graphql/app/interfaces"
	"github.com/meles-z/golang-graphql/configs"
	"github.com/vektah/gqlparser/v2/ast"
)

func graphqlHandler(resolver *interfaces.Resolver) http.Handler {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	gormDB, err := db.InitDB(cfg.DB)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	userRepo := persistence.NewUserRepository(gormDB)
	movieRepo := persistence.NewMovieRepository(gormDB)

	resolver := &interfaces.Resolver{
		UserRepo:  userRepo,
		MovieRepo: movieRepo,
	}

	// 💡 Fix: use the returned handler here
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graphqlHandler(resolver))

	log.Println("🚀 Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
