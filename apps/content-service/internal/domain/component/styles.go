package component

var stylesBreakpoint = map[string]struct{}{
	"base":          {},
	"mobile":        {},
	"tablet":        {},
	"desktop":       {},
	"large_desktop": {},
}

type StyleBreakpoint string

func NewStyleBreakpoint(s string) (StyleBreakpoint, error) {
	if _, ok := stylesBreakpoint[s]; !ok {
		return "", ErrInvalidStyleBreakpoint
	}
	return StyleBreakpoint(s), nil
}
