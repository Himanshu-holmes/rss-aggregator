package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/himanshu-holmes/rss-aggregator/internal/database"
)




func (apiCfg apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter,req *http.Request, user database.User){
	
	feedFollowsIDStr := chi.URLParam(req,"feedFollowID");
	feedFollowID,err:=uuid.Parse(feedFollowsIDStr);

    if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not parse feed_follow Id: %v",err));
		return
	 }
	 err = apiCfg.DB.DeleteFeedFollows(req.Context(),database.DeleteFeedFollowsParams{
		ID: feedFollowID,
		UserID: user.ID,
	 });
	 if err != nil {
		respondWithError(w,400,fmt.Sprintf("Could not delete feed_follow : %v",err));
		return
	 }
	 respondWithJson(w,200,struct{}{})

}  