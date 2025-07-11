package page

import (
	"errors"
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

func (p *Page) EditComponent(componentID component.ComponentID, updatedComponent *component.Component) error {
	foundComponent, isFound := p.findComponentById(componentID)
	if !isFound {
		return errors.New("Component not found")
	}
	foundComponent.Update(updatedComponent)
	return nil
}

func (p *Page) findComponentById(id component.ComponentID) (*component.Component, bool) {
	for i := range p.Components {
		if p.Components[i].ComponentID == id {
			return p.Components[i], true
		}
	}
	return nil, false
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
