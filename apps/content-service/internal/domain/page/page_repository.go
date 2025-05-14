package page

type PageRepository interface {
	Insert(page *Page) error
	Update(page *Page) error
}
