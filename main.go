package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	// "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world");
	godotenv.Load(".env")
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		log.Fatal("PORT is not found in the env")
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

   srv := &http.Server{
	  Handler: router,
	  Addr: ":"+ portEnv,
   }
log.Printf("server starting on %v",portEnv)
 err :=  srv.ListenAndServe()
 if err !=nil {
	log.Fatal((err))
 }


	fmt.Println("Port:",portEnv)
}
