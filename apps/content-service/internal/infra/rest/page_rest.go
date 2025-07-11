package rest

import (
	"encoding/json"
	"net/http"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type PageRest struct {
	pageSvc page.PageService
}

func NewPageRest(sm *http.ServeMux, pageService page.PageService) {
	pageRest := &PageRest{
		pageSvc: pageService,
	}

	sm.HandleFunc("POST /v1/pages", pageRest.createPage)
	sm.HandleFunc("GET /v1/pages", pageRest.getPages)
	sm.HandleFunc("POST /v1/pages/{id}/publish", pageRest.publishPage)
}

func (pr *PageRest) createPage(w http.ResponseWriter, r *http.Request) {
	var dto createPageDTO
	if err := readJSON(w, r, &dto); err != nil {
		return
	}

	page, err := pr.pageSvc.CreateDraftPage(dto.Title)

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

	pages, err := pr.pageSvc.GetPages(filter)
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

	page, err := pr.pageSvc.GetPageById(pageID)
	if err != nil {
		writeErr(w, err)
		return
	}
	if page == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := pr.pageSvc.PublishPage(page); err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, http.StatusOK, page)
}
