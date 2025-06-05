package component

type GlobalComponentRepository interface {
	Insert(c *GlobalComponent) error
	Update(c *GlobalComponent) error
}
