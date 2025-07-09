package services

import (
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type PageServiceImpl struct {
	page.PageService
	page.PageRepository
}

func (ps *PageServiceImpl) CreateDraftPage() (*page.Page, error) {
	p := page.NewDraft()

	if err := ps.Insert(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (ps *PageServiceImpl) PublishPage(p *page.Page) error {
	p.MarkAsPublished()

	if err := ps.Update(p); err != nil {
		return err
	}

	return nil
}

func (ps *PageServiceImpl) GetPages(filter page.GetPages) ([]*page.Page, error) {
	pages, err := ps.PageRepository.FindByTitle(filter.Title)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func (ps *PageServiceImpl) GetPageById(id page.PageID) (*page.Page, error) {
	page, err := ps.PageRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (ps *PageServiceImpl) AddComponent(p *page.Page, comp *component.Component) error {
	p.AddComponent(comp)

	if err := ps.Update(p); err != nil {
		return err
	}

	return nil
}
