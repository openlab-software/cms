package rest

type addComponentDTO struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}
