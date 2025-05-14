package page

import (
	"time"

	"github.com/patrickdevbr-portfolio/cms/libs/go-common/publicid"
)

const (
	DRAFT     = "DRAFT"
	PUBLISHED = "PUBLISHED"
	INACTIVE  = "INACTIVE"
)

type PageID publicid.PublicID

type Page struct {
	ID           PageID
	Title        string
	Status       string
	CreatedAt    time.Time
	PublisheddAt *time.Time
}

func (p *Page) Publish() {
	publishedAt := time.Now()

	p.Status = PUBLISHED
	p.PublisheddAt = &publishedAt
}

func NewDraft() *Page {
	return &Page{
		ID:           PageID(publicid.New()),
		Title:        "Draft",
		Status:       DRAFT,
		CreatedAt:    time.Now(),
		PublisheddAt: nil,
	}
}
