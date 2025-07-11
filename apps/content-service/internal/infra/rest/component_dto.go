package rest

import "github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"

type editComponentDTO struct {
	Type   component.ComponentType           `json:"type"`
	Data   map[string]any                    `json:"data"`
	Styles map[component.StyleBreakpoint]any `json:"styles"`
}
