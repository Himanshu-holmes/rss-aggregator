package main

import "net/http"

func handlerErr(w http.ResponseWriter, req *http.Request){
	respondWithJson(w,400,"Something went wrong")
}