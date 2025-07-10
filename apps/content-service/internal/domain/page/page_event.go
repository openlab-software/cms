package page

import (
	"time"
)

const (
	PageDraftedEvent   = "page.drafted"
	PagePublishedEvent = "page.published"
)

type PageDraftedPayload struct {
	ID    PageID `json:"page_ id"`
	Title string `json:"title"`
}

type PagePublishedPayload struct {
	ID          PageID    `json:"page_id"`
	PublishedAt time.Time `json:"publishedAt"`
}
