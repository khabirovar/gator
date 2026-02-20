package main

import (
	"fmt"
	"context"
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
