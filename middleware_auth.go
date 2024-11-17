package main

import (
	"fmt"
	"net/http"

	"github.com/himanshu-holmes/rss-aggregator/internal/auth"
	"github.com/himanshu-holmes/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter,*http.Request,database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request){
		apiKey,err := auth.GetAPIKey(req.Header);
	 
		if err != nil {
		   respondWithError(w,403,fmt.Sprintf("Auth Error :%v",err))
		   return 
		}
   
	   user ,err := cfg.DB.GetUserByAPIKey(req.Context(),apiKey)
   
	   if err != nil {
		   respondWithError(w,403,fmt.Sprintf("Couldn't get user :%v",err))
		   return 
		};
   
		
		handler(w,req,user)
	}
}