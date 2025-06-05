package component

type ComponentService interface {
	SetAsGlobal(c *Component) (GlobalComponent, error)
}
