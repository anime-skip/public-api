// This script was created apart of https://github.com/anime-skip/player/issues/246 and it:
//
// 1. Goes through all the shows in the database
// 2. Looks them up by name in anilist
// 3. Adds the Anilist site URL as an external link for the show IF a match is found

package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/http"
	"anime-skip.com/public-api/internal/postgres"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env.local")
	// godotenv.Load(".env.prod")
}

func main() {
	ctx := context.Background()
	anilist := http.NewAnilistService()
	db := postgres.Open(
		os.Getenv("DATABASE_URL"),
		os.Getenv("DATABASE_DISABLE_SSL") == "true",
		nil,
		false,
	)
	showService := postgres.NewShowService(db, anilist)
	externalLinks := postgres.NewExternalLinkService(db)

	println("Getting all shows...")
	shows, err := showService.List(ctx, internal.ShowsFilter{})
	checkErr(err)

	showCount := len(shows)
	println(showCount)

	println("Looking for Anilist matches:")
	for i, show := range shows {
		time.Sleep(time.Minute / 90.0 /* req */) // Prevent rate limiting
		link, err := anilist.FindLink(show.Name)
		if link == nil {
			fmt.Printf("(%d/%d) %s - NO URL\n", i+1, showCount, show.Name)
		} else {
			_, err = externalLinks.Create(ctx, internal.ExternalLink{
				URL:    *link,
				ShowID: show.ID,
			})
			if err != nil && !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint \"external_links_pkey\"") {
				checkErr(err)
			}
			fmt.Printf("(%d/%d) %s - %s\n", i+1, showCount, show.Name, *link)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		println("Error: " + err.Error())
		os.Exit(1)
	}
}
