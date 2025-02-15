package types

type ContentData interface {
	GetURL() string
	GetType() string
}

type Content struct {
	URL  string
	Type string
}

func (c *Content) GetURL() string {
	return c.URL
}

func (c *Content) GetType() string {
	return c.Type
}

type YoutubeContent struct {
	Content
	Title  string // Video title
	Author string // Channel name
}
