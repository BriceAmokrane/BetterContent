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
	// Extraction de l'ID de la vidéo depuis l'URL
	videoID, err := extractVideoID(url)
	if err != nil {
		return nil, err
	}

	// Appel à l'API YouTube
	call := s.service.Videos.List([]string{"snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error calling YouTube API: %v", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("no video found with ID: %s", videoID)
	}

	video := response.Items[0]
	content := &types.YoutubeContent{
		Content: types.Content{
			URL:  url,
			Type: "youtube",
		},
		Title:       video.Snippet.Title,
		Author:      video.Snippet.ChannelTitle,
		Description: video.Snippet.Description,
	}

	return content, nil
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
