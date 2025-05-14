package page

import (
	"time"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/audit"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/publicid"
)

const (
	DRAFT     = "DRAFT"
	PUBLISHED = "PUBLISHED"
	INACTIVE  = "INACTIVE"
)

type PageID publicid.PublicID

type Page struct {
	audit.Audit
	ID          PageID
	Title       string
	Status      string
	Components  []component.Component
	PublishedAt *time.Time
}

func (p *Page) MarkAsPublished() {
	publishedAt := time.Now()

	p.Status = PUBLISHED
	p.PublishedAt = &publishedAt
}

func NewDraft() *Page {
	return &Page{
		ID:          PageID(publicid.New()),
		Title:       "Draft",
		Status:      DRAFT,
		PublishedAt: nil,
		Audit: audit.Audit{
			CreatedAt: time.Now(),
		},
	}
}
