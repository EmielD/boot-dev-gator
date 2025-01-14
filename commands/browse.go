package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
)

func HandlerBrowse(s *types.State, cmd types.Command, user database.User) error {
	limit := 2
	if len(cmd.Arguments) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Arguments[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {

		publishedAtTime := ""
		if post.PublishedAt.Valid {
			publishedAtTime = fmt.Sprintf("%v from", post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf("%s from %s\n", publishedAtTime, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
