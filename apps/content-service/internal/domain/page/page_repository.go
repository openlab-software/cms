package page

type PageRepository interface {
	Insert(page *Page) error
	Update(page *Page) error
	FindByTitle(title string) ([]*Page, error)
	FindById(id PageID) (*Page, error)
}
