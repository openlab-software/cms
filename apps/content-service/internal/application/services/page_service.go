package services

import "github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"

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
	p.Publish()

	if err := ps.Update(p); err != nil {
		return err
	}

	return nil
}
