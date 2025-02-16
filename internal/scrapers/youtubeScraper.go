package scrapers

import "BetterContent/internal/types"

type YoutubeScraper struct{}

func NewYoutubeScraper() ContentScraper {
	return &YoutubeScraper{}
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
