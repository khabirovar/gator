package main

import (
	"fmt"
	"context"
	"errors"
	"time"
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

	for _, article := range rssFeed.Channel.Item {
		fmt. Printf("  * %s\n", article.Title)
	}
	return nil
}
