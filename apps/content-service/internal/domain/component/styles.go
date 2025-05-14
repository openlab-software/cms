package component

var stylesBreakpoint = map[string]struct{}{
	"TEXT":      {},
	"IMAGE":     {},
	"BUTTON":    {},
	"CONTAINER": {},
	"COLUMN":    {},
}

type StyleBreakpoint string

func NewStyleBreakpoint(s string) (StyleBreakpoint, error) {
	if _, ok := stylesBreakpoint[s]; !ok {
		return "", ErrInvalidStyleBreakpoint
	}
	return StyleBreakpoint(s), nil
}
