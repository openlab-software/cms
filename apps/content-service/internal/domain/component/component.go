package component

import (
	"errors"
	"time"

	"github.com/patrickdevbr-portfolio/cms/libs/go-common/audit"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/publicid"
)

var ErrInvalidStyleBreakpoint = errors.New("invalid style breakpoint")

type ComponentID = publicid.PublicID

type Component struct {
	audit.Audit
	ComponentID ComponentID             `bson:"component_id" json:"component_id"`
	GlobalID    *ComponentID            `bson:"global_id" json:"global_id"`
	Type        ComponentType           `bson:"type" json:"type"`
	Data        map[string]any          `bson:"data" json:"data"`
	Styles      map[StyleBreakpoint]any `bson:"styles" json:"styles"`
}

func NewComponent(Type ComponentType, Data map[string]any, Styles map[StyleBreakpoint]any) *Component {
	return &Component{
		ComponentID: ComponentID(publicid.New("component")),
		Type:        Type,
		Data:        Data,
		Styles:      Styles,
		Audit: audit.Audit{
			CreatedAt: time.Now(),
		},
	}
}

func (c Component) IsGlobal() bool {
	return c.GlobalID == nil
}

func ParseComponentID(s string) (ComponentID, error) {
	publicID, err := publicid.Parse("component", s)

	return ComponentID(publicID), err
}
