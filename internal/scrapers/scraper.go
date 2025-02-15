package scrapers

import (
	"BetterContent/internal/types"
)

type ContentScraper interface {
	Scrape(url string) (types.ContentData, error)
}
