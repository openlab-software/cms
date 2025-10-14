package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/component"
	"github.com/patrickdevbr-portfolio/cms/apps/content-service/internal/domain/page"
)

type ComponentRest struct {
	pageSvc page.PageService
}

func NewComponentRest(r *mux.Router, pageSvc page.PageService) {
	componentRest := &ComponentRest{
		pageSvc: pageSvc,
	}

	componentRouter := r.PathPrefix("/pages/{pageID}/components").Subrouter()

	componentRouter.HandleFunc("/{componentID}", componentRest.editComponent).Methods("PATCH")
	componentRouter.HandleFunc("", componentRest.addComponent).Methods("POST")
	componentRouter.HandleFunc("", componentRest.getComponentById).Methods("GET")
}

func (cr *ComponentRest) editComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pageID, err := page.ParsePageID(vars["pageID"])
	if err != nil {
		writeErr(w, err)
		return
	}
	componentID, err := component.ParseComponentID(vars["componentID"])
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
	vars := mux.Vars(r)

	pageID, err := page.ParsePageID(vars["pageID"])
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

func (cr *ComponentRest) getComponentById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
