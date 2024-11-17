package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/himanshu-holmes/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`;
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Name      string      `json:"name"`
	APIKey    string       `json:"api_key"`
}

func databaseUserToUser(dbUser database.User)User{
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`;
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Name      string      `json:"name"`
	Url    string       `json:"url"`
	UserId uuid.UUID     `json:"user_id"`
 }

 
func databaseFeedToFeed(dbFeed database.Feed)Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserId: dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed)[]Feed{
	feeds:= []Feed{};
	for _,dbFeed := range dbFeeds {
	feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
  }

  type FeedFollows struct {
	ID        uuid.UUID `json:"id"`;
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UserId uuid.UUID     `json:"user_id"`
	FeeID uuid.UUID       `json:"feed_id"`
 }

 func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollows{
	return FeedFollows{
		ID: dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserId: dbFeedFollow.UserID,
		FeeID: dbFeedFollow.FeedID,
	}
 }

 func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow)[]FeedFollows{
	feedFollows:= []FeedFollows{};
	for _,dbFeedFollow := range dbFeedFollows {
	feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
  }