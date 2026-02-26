package main

import (
	"context"
	"fmt"
	"github.com/khabirovar/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) > 0 {
		limitConv, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(limitConv)
	}

	paramsPosts := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), paramsPosts)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println("=======================================================")
		fmt.Printf("Title:\t%s\n", post.Title)
		fmt.Printf("Url:\t%s\n", post.Url)
		fmt.Printf("Description:\t%s\n", post.Description.String)
	}
	return nil
}
