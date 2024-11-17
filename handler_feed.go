package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	
	"github.com/himanshu-holmes/rss-aggregator/internal/database"
)

func (apiCfg apiConfig) handlerCreateFeed(w http.ResponseWriter,r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url  string  `json:"url"`
	}
	decoder := json.NewDecoder(r.Body);

	params := parameters{};

	err := decoder.Decode(&params);

	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %v",err))
		return
	}
	feed,err := apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:  params.Name,
		Url:params.Url,
		UserID: user.ID,
	 });





	 if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not create user : %v",err));
		return
	 }
	respondWithJson(w, 200,databaseFeedToFeed(feed))
}


func (apiCfg apiConfig) handlerGetFeed(w http.ResponseWriter,req *http.Request){
	
	feeds,err := apiCfg.DB.GetFeeds(req.Context());





	 if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not create user : %v",err));
		return
	 }
	respondWithJson(w, 200,databaseFeedsToFeeds(feeds))
	
}  