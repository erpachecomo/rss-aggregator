package main

import (
	"context"
	"fmt"
	"rss-aggregator/internal/rss"
)

func handleAgg(state *state, cmd command) error {
	rssUrl := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), rssUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Channel title: %v\n Channel description: %v\n Channel link: %v", feed.Channel.Title, feed.Channel.Description, feed.Channel.Link)
	for index, item := range feed.Channel.Item {
		fmt.Printf("%v. Item title: %v\n Item description: %v\n Item link: %v\n Item Pub Date: %v", index, item.Title, item.Description, item.Link, item.PubDate)

	}
	return nil
}
