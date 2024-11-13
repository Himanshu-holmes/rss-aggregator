package main

import "net/http"

func handlerReadiness(w http.ResponseWriter,req *http.Request){
	respondWithJson(w,200,struct{}{})
}