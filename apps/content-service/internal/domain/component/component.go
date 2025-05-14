package component

import (
	"errors"

	"github.com/patrickdevbr-portfolio/cms/libs/go-common/audit"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/publicid"
)

var ErrInvalidStyleBreakpoint = errors.New("invalid style breakpoint")

type ComponentID = publicid.PublicID

type Component struct {
	audit.Audit
	ID       ComponentID
	GlobalID *ComponentID
	Type     ComponentType
	Data     map[StyleBreakpoint]any
	Styles   map[StyleBreakpoint]any
}

func (c Component) IsGlobal() bool {
	return c.GlobalID == nil
}
