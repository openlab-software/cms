package rest

import (
	"time"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/audit"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/publicid"
)

type createPageDTO struct {
	Title string `json:"title"`
}

type addComponentDTO struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type pageDTO struct {
	audit.Audit
	PageID      string          `json:"page_id"`
	Title       string          `json:"title"`
	Status      string          `json:"status"`
	Components  *[]componentDTO `json:"components"`
	PublishedAt *time.Time      `json:"published_at"`
}

type componentDTO struct {
	audit.Audit
	ComponentID string                            `json:"component_id"`
	GlobalID    *string                           `json:"global_id"`
	Type        component.ComponentType           `json:"type"`
	Data        map[string]any                    `json:"data"`
	Styles      map[component.StyleBreakpoint]any `json:"styles"`
}

func toComponentDTO(c *component.Component) componentDTO {
	return componentDTO{
		ComponentID: publicid.PublicID(c.ComponentID).ToPublic(),
		Audit:       c.Audit,
		GlobalID:    (*string)(c.GlobalID),
		Type:        c.Type,
		Data:        c.Data,
		Styles:      c.Styles,
	}
}

func toComponentsDTO(comps []*component.Component) *[]componentDTO {
	dtos := make([]componentDTO, len(comps))
	for i, c := range comps {
		dtos[i] = toComponentDTO(c)
	}
	return &dtos
}

func toPageDTO(page *page.Page) pageDTO {
	return pageDTO{
		PageID:      publicid.PublicID(page.PageID).ToPublic(),
		Audit:       page.Audit,
		Title:       page.Title,
		Status:      page.Status,
		Components:  toComponentsDTO(page.Components),
		PublishedAt: page.PublishedAt,
	}
}

func toPagesDTO(pages []*page.Page) *[]pageDTO {
	dtos := make([]pageDTO, len(pages))
	for i, p := range pages {
		dtos[i] = toPageDTO(p)
	}
	return &dtos
}
