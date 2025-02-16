package scrapers

import (
	"BetterContent/internal/types"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"regexp"
)

type YoutubeScraper struct {
	service *youtube.Service
}

func NewYoutubeScraper(apiKey string) (*YoutubeScraper, error) {
	ctx := context.Background()
	service, err := youtube.NewService(
		ctx,
		option.WithAPIKey(apiKey),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create YouTube service: %v", err)
	}

	return &YoutubeScraper{service: service}, nil
}

func (s *YoutubeScraper) Scrape(url string) (types.ContentData, error) {
	return &types.YoutubeContent{
		Content: types.Content{
			URL:  url,
			Type: "youtube",
		},
		Title:  "Dummy title",
		Author: "Dummy author",
	}, nil
}

func extractVideoID(url string) (string, error) {
	var videoID string

	// Patterns possibles pour les URLs YouTube
	patterns := map[string]*regexp.Regexp{
		"standard":  regexp.MustCompile(`(?:youtube\.com/watch\?v=|youtu.be/)([^&?/]+)`),
		"shortened": regexp.MustCompile(`youtu\.be/([^?/]+)`),
		"embed":     regexp.MustCompile(`youtube\.com/embed/([^?/]+)`),
	}

	for _, pattern := range patterns {
		matches := pattern.FindStringSubmatch(url)
		if len(matches) > 1 {
			videoID = matches[1]
			break
		}
	}

	if videoID == "" {
		return "", fmt.Errorf("could not extract video ID from URL: %s", url)
	}

	return videoID, nil
}
