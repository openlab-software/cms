package rest

import (
	"net/http"

	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type ComponentRest struct {
	pageSvc page.PageService
}

func NewComponentRest(sm *http.ServeMux, pageSvc page.PageService) {
	componentRest := &ComponentRest{
		pageSvc: pageSvc,
	}

	sm.HandleFunc("PATCH /v1/pages/{pageID}/components/{componentID}", componentRest.editComponent)
	sm.HandleFunc("POST /v1/pages/{id}/components", componentRest.addComponent)
}

func (cr *ComponentRest) editComponent(w http.ResponseWriter, r *http.Request) {
	pageID, err := page.ParsePageID(r.PathValue("pageID"))
	if err != nil {
		writeErr(w, err)
		return
	}
	componentID, err := component.ParseComponentID(r.PathValue("componentID"))
	if err != nil {
		writeErr(w, err)
		return
	}

	page, err := cr.pageSvc.GetPageById(pageID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var dto editComponentDTO
	if err := readJSON(w, r, &dto); err != nil {
		return
	}

	err = cr.pageSvc.EditComponent(page, componentID, &component.Component{Data: dto.Data, Type: dto.Type, Styles: dto.Styles})

	if err != nil {
		writeErr(w, err)
		return
	}

	writeJSON(w, http.StatusOK, page)
}

func (cr *ComponentRest) addComponent(w http.ResponseWriter, r *http.Request) {
	pageID, err := page.ParsePageID(r.PathValue("id"))
	if err != nil {
		writeErr(w, err)
		return
	}

	page, err := cr.pageSvc.GetPageById(pageID)
	if err != nil {
		writeErr(w, err)
		return
	}

	var dto addComponentDTO
	if err := readJSON(w, r, &dto); err != nil {
		return
	}

	compType, err := component.NewComponentType(dto.Type)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	newComponent := component.NewComponent(compType, dto.Data, nil)

	if err := cr.pageSvc.AddComponent(page, newComponent); err != nil {
		writeErr(w, err)
		return
	}
	writeJSON(w, http.StatusOK, page)
}
