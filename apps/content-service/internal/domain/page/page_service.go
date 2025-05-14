package page

type PageService interface {
	CreateDraftPage() (*Page, error)
	PublishPage(p *Page) error
}
