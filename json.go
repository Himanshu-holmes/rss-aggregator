package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func repondWithJson(w http.ResponseWriter,code int,payload interface{}){
	data,err := json.Marshal(payload);
	if err != nil {
		log.Printf("Failed to Marshal json response %v",payload);
		w.WriteHeader(500);
		return
	}
	w.Header().Add("content-type","application/json")
	w.WriteHeader(200);
	w.Write(data)
}