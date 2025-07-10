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
	PageID      PageID                 `bson:"page_id" json:"page_id"`
	Title       string                 `bson:"title" json:"title"`
	Status      string                 `bson:"status" json:"status"`
	Components  []*component.Component `bson:"components" json:"components"`
	PublishedAt *time.Time             `bson:"published_at" json:"published_at"`
}

func (p *Page) MarkAsPublished() {
	publishedAt := time.Now()

	p.Status = PUBLISHED
	p.PublishedAt = &publishedAt
}

func (p *Page) AddComponent(comp *component.Component) {
	p.Components = append(p.Components, comp)
}

func NewDraft(title string) *Page {
	return &Page{
		PageID:      PageID(publicid.New("page")),
		Title:       title,
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
