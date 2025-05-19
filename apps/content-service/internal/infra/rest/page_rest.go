package rest

import (
	"encoding/json"
	"fmt"
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
	sm.HandleFunc("POST /v1/pages/{id}/publish", pageRest.publishPage)
}

func (pr *PageRest) createPage(w http.ResponseWriter, r *http.Request) {
	page, err := pr.PageService.CreateDraftPage()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(page)
}

func (pr *PageRest) getPages(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := page.GetPages{
		Title: query.Get("title"),
	}

	pages, err := pr.PageService.GetPages(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pages)
}

func (pr *PageRest) publishPage(w http.ResponseWriter, r *http.Request) {
	pageID, err := page.ParsePageID(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
		return
	}

	page, err := pr.PageService.GetPageById(pageID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
		return
	}

	if err := pr.PageService.PublishPage(page); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]any{"err": err})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(page)
}
