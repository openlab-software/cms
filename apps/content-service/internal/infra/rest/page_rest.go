package rest

import (
	"encoding/json"
	"net/http"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
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
	sm.HandleFunc("POST /v1/pages/{id}/components", pageRest.addComponent)
}

func (pr *PageRest) createPage(w http.ResponseWriter, r *http.Request) {
	page, err := pr.PageService.CreateDraftPage()

	if err != nil {
		writeErr(w, err)
		return
	}

	writeJSON(w, http.StatusOK, page)
}

func (pr *PageRest) getPages(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filter := page.GetPages{
		Title: query.Get("title"),
	}

	pages, err := pr.PageService.GetPages(filter)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pages)
}

func (pr *PageRest) publishPage(w http.ResponseWriter, r *http.Request) {
	pageID, err := page.ParsePageID(r.PathValue("id"))
	if err != nil {
		writeErr(w, err)
		return
	}

	page, err := pr.PageService.GetPageById(pageID)
	if err != nil {
		writeErr(w, err)
		return
	}

	if err := pr.PageService.PublishPage(page); err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, http.StatusOK, page)
}

func (pr *PageRest) addComponent(w http.ResponseWriter, r *http.Request) {
	pageID, err := page.ParsePageID(r.PathValue("id"))
	if err != nil {
		writeErr(w, err)
		return
	}

	page, err := pr.PageService.GetPageById(pageID)
	if err != nil {
		writeErr(w, err)
		return
	}

	var dto addComponentDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	compType, err := component.NewComponentType(dto.Type)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	newComponent := component.NewComponent(compType, dto.Data, nil)

	if err := pr.PageService.AddComponent(page, newComponent); err != nil {
		writeErr(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(page)
}
