package handlers

import (
	"BetterContent/internal/scrapers"
	"BetterContent/internal/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

type ContentHandler struct {
	urlValidator validators.URLValidator
	scrapers     map[string]scrapers.ContentScraper
}

func NewContentHandler(validator validators.URLValidator) *ContentHandler {
	youtubeScraper, _ := scrapers.NewYoutubeScraper(os.Getenv("YOUTUBE_API_KEY"))
	return &ContentHandler{
		urlValidator: validator,
		scrapers: map[string]scrapers.ContentScraper{
			"youtube.com": youtubeScraper,
			"youtu.be":    youtubeScraper,
		},
	}
}

func (h *ContentHandler) HandleAddContentLink(c *gin.Context) {
	contentAddr := c.PostForm("contentAddr")

	// Validation de l'URL
	if err := h.urlValidator.ValidateURL(contentAddr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Vérification que l'URL est accessible
	if err := h.urlValidator.IsAlive(contentAddr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extraction du domaine
	parsedURL, err := url.Parse(contentAddr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}

	// Sélection du scraper approprié
	scraper, found := h.getScraper(parsedURL.Host)
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported domain"})
		return
	}

	// Scraping du contenu
	content, err := scraper.Scrape(contentAddr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *ContentHandler) getScraper(domain string) (scrapers.ContentScraper, bool) {
	for supportedDomain, scraper := range h.scrapers {
		if strings.Contains(domain, supportedDomain) {
			return scraper, true
		}
	}
	return nil, false
}
