package services

import (
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/event"
)

type PageServiceImpl struct {
	page.PageService
	repository     page.PageRepository
	eventPublisher event.Publisher
}

func NewPageService(repo page.PageRepository, eventPublisher event.Publisher) page.PageService {
	return &PageServiceImpl{
		repository:     repo,
		eventPublisher: eventPublisher,
	}
}

func (ps *PageServiceImpl) CreateDraftPage(title string) (*page.Page, error) {
	p := page.NewDraft(title)

	if err := ps.repository.Insert(p); err != nil {
		return nil, err
	}

	draftedEvent := event.NewEvent(page.PageDraftedEvent, page.PageDraftedPayload{ID: p.PageID, Title: p.Title})

	if err := ps.eventPublisher.Publish(draftedEvent); err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *PageServiceImpl) PublishPage(p *page.Page) error {
	p.MarkAsPublished()

	if err := ps.repository.Update(p); err != nil {
		return err
	}

	publishedEvent := event.NewEvent(page.PagePublishedEvent, page.PagePublishedPayload{ID: p.PageID, PublishedAt: *p.PublishedAt})

	if err := ps.eventPublisher.Publish(publishedEvent); err != nil {
		return err
	}
	return nil
}

func (ps *PageServiceImpl) GetPages(filter page.GetPages) ([]*page.Page, error) {
	pages, err := ps.repository.FindByTitle(filter.Title)
	if err != nil {
		return nil, err
	}
	return pages, nil
}

func (ps *PageServiceImpl) GetPageById(id page.PageID) (*page.Page, error) {
	page, err := ps.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (ps *PageServiceImpl) AddComponent(p *page.Page, comp *component.Component) error {
	p.AddComponent(comp)

	if err := ps.repository.Update(p); err != nil {
		return err
	}

	return nil
}
