package main

import (
	"fmt"
	"context"
	"errors"
	"time"
	"github.com/google/uuid"
	"github.com/khabirovar/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err 
	}
	fmt.Printf("RSSFeed: %#v\n", rssFeed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("HandlerAddFeed expects two arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	
	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %#v\n", feed)

	return nil
}

func handlerFeeds(s* state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("Feed: %s URL: %s CreatedBy: %s\n", feed.Feed, feed.Url, feed.User)
	}
	return nil
}
