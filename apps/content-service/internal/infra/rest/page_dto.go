package rest

type createPageDTO struct {
	Title string `json:"title"`
}

type addComponentDTO struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}
