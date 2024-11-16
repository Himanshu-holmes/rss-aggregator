package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/himanshu-holmes/rss-aggregator/internal/database"

	// "github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}
func main() {
	fmt.Println("hello world");
	godotenv.Load(".env")
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		log.Fatal("PORT is not found in the env")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl== "" {
		log.Fatal("DB_URL is not found in the env")
	}

    conn,err := sql.Open("postgres",dbUrl);

	if err != nil {
		log.Fatal("Can't connect to database",err)
	}

	

	apiCfg := apiConfig{
		DB: database.New(conn),
	}


   router := chi.NewRouter();
   router.Use(cors.Handler(cors.Options{
	AllowedOrigins: []string{"https://*","http://*"},
	AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
	AllowedHeaders: []string{"*"},
	ExposedHeaders: []string{"Link"},
	AllowCredentials: false,
	MaxAge: 300,
   }))

   v1Router := chi.NewRouter();
   v1Router.Get("/healthz",handlerReadiness);
   v1Router.Get("/err",handlerErr)
   v1Router.Post("/users",apiCfg.handlerCreatUser)
   router.Mount("/v1",v1Router);

   srv := &http.Server{
	  Handler: router,
	  Addr: ":"+ portEnv,
   }
  log.Printf("server starting on %v",portEnv)
   err =  srv.ListenAndServe()
 if err !=nil {
	log.Fatal((err))
 }


	fmt.Println("Port:",portEnv)
}
