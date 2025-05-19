package page

type PageService interface {
	CreateDraftPage() (*Page, error)
	PublishPage(p *Page) error
	GetPages(filter GetPages) ([]*Page, error)
	GetPageById(id PageID) (*Page, error)
}

type GetPages struct {
	Title string
}
