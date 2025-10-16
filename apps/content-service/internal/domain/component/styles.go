package component

import (
	"fmt"
)

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
		return "", fmt.Errorf("'%s'is a invalid style breakpoint", s)
	}
	return StyleBreakpoint(s), nil
}

func ValidateStyles(styles map[StyleBreakpoint]any) error {
	for s := range styles {
		_, err := NewStyleBreakpoint(string(s))
		if err != nil {
			return err
		}
	}
	return nil
}
