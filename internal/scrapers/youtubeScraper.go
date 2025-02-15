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
