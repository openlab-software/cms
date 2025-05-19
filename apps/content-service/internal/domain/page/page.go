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
	PageID      PageID `bson:"page_id" json:"page_id"`
	Title       string
	Status      string
	Components  []component.Component
	PublishedAt *time.Time `bson:"published_at" json:"published_at"`
}

func (p *Page) MarkAsPublished() {
	publishedAt := time.Now()

	p.Status = PUBLISHED
	p.PublishedAt = &publishedAt
}

func NewDraft() *Page {
	return &Page{
		PageID:      PageID(publicid.New("page")),
		Title:       "Draft",
		Status:      DRAFT,
		PublishedAt: nil,
		Audit: audit.Audit{
			CreatedAt: time.Now(),
		},
	}
}

func ParsePageID(s string) (PageID, error) {
	publicID, err := publicid.Parse("page", s)

	return PageID(publicID), err
}
