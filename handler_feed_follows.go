package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	
	"github.com/himanshu-holmes/rss-aggregator/internal/database"
)

func (apiCfg apiConfig) handlerCreateFeedFollow(w http.ResponseWriter,r *http.Request, user database.User){
	type parameters struct {
		FeeID uuid.UUID `json:"feed_id"`
		
	}
	decoder := json.NewDecoder(r.Body);

	params := parameters{};

	err := decoder.Decode(&params);

	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %v",err))
		return
	}
	feedFollow,err := apiCfg.DB.CreateFeedFollows(r.Context(),database.CreateFeedFollowsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
        FeedID: params.FeeID,
		UserID: user.ID,
	 });





	 if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not create user : %v",err));
		return
	 }
	respondWithJson(w, 200,databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg apiConfig) handlerGetFeedFollows(w http.ResponseWriter,req *http.Request, user database.User){
	fmt.Printf("user id ::%v",user.ID)
	feedFollows,err := apiCfg.DB.GetFeedFollows(req.Context(),user.ID);

	 if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not create user : %v",err));
		return
	 }
	respondWithJson(w, 200,databaseFeedFollowsToFeedFollows(feedFollows))
	
}  