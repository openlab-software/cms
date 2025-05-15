package rest

import (
	"encoding/json"
	"net/http"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type PageRest struct {
	page.PageService
}

func NewPageRest(sm *http.ServeMux, pageService page.PageService) {
	pageRest := &PageRest{
		PageService: pageService,
	}

	sm.HandleFunc("POST /v1/pages", pageRest.createPage)
	sm.HandleFunc("GET /v1/pages", pageRest.getPages)
}

func (pr *PageRest) createPage(w http.ResponseWriter, r *http.Request) {
	page, err := pr.PageService.CreateDraftPage()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(page)
}

func (pr *PageRest) getPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(page.NewDraft())
}
