package component

import "errors"

type ComponentType string

var ErrInvalidComponentType = errors.New("invalid component type")

var componentTypes = map[string]struct{}{
	"base":          {},
	"mobile":        {},
	"tablet":        {},
	"desktop":       {},
	"large_desktop": {},
}

func NewComponentType(s string) (ComponentType, error) {
	if _, ok := componentTypes[s]; !ok {
		return "", ErrInvalidComponentType
	}
	return ComponentType(s), nil
}

func (ct ComponentType) IsValid() bool {
	_, ok := componentTypes[string(ct)]
	return ok
}
