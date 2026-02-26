package main

import (
	"fmt"
	"context"
	"errors"
	"time"
	"github.com/lib/pq"
	"github.com/khabirovar/gator/internal/database"
	"database/sql"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("HandlerAgg expects one argument")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenReqs)
	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])
	for ;; <- ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Printf("Couldn't collect feed: %v\n", err) 
		}
	}
	return nil 
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Println("Posts added to database:")
	for _, article := range rssFeed.Channel.Item {
		if article.Link == "" {
			continue
		}
		pubDate, err := time.Parse(time.RFC1123Z, article.PubDate)
		nullPubDate := sql.NullTime{Time: pubDate, Valid: err == nil}
		nullDescription := sql.NullString{String: article.Description, Valid: article.Description != ""}
		createPostParam := database.CreatePostParams{
			Title: article.Title,
			Url: article.Link,
			Description:  nullDescription,
			PublishedAt:  nullPubDate,
			FeedID: feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), createPostParam)
		if err != nil {
			var pgError *pq.Error
			if errors.As(err, &pgError) && pgError.Code == "23505" {
				continue
			}
			return err 
		}
		fmt.Printf("  * %s\n", article.Title)
	}
	return nil
}
