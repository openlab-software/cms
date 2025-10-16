package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type PageRest struct {
	pageSvc page.PageService
}

func NewPageRest(r *mux.Router, pageService page.PageService) {
	pageRest := &PageRest{
		pageSvc: pageService,
	}

	pageRouter := r.PathPrefix("/pages").Subrouter()

	pageRouter.HandleFunc("", pageRest.createPage).Methods("POST")
	pageRouter.HandleFunc("", pageRest.getPages).Methods("GET")
	pageRouter.HandleFunc("/{pageID}/publish", pageRest.publishPage).Methods("POST")
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

	writeJSON(w, http.StatusOK, toPageDTO(page))
}

// @Summary Edita um componente existente
// @Description Atualiza parcialmente um componente de uma página
// @Tags components
// @Param pageID path string true "ID da página"
// @Param componentID path string true "ID do componente"
// @Accept json
// @Produce json
// @Success 200 {string} string "Componente atualizado"
// @Failure 400 {string} string "Requisição inválida"
// @Router /pages/{pageID}/components/{componentID} [patch]
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

	json.NewEncoder(w).Encode(toPagesDTO(pages))
}

func (pr *PageRest) publishPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pageID, err := page.ParsePageID(vars["pageID"])
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
	writeJSON(w, http.StatusOK, toPageDTO(page))
}
