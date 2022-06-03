package server

import (
	"net/http"

	"github.com/frain-dev/convoy/server/models"
	"github.com/frain-dev/convoy/util"
	"github.com/go-chi/render"
)

// LoadConfiguration
// @Summary Fetch configuration
// @Description This endpoint fetches configuration
// @Tags Source
// @Accept  json
// @Produce  json
// @Success 200 {object} serverResponse{data=pagedResponse{content=[]datastore.Configuration}}
// @Failure 400,401,500 {object} serverResponse{data=Stub}
// @Security ApiKeyAuth
// @Router /configuration [get]
func (a *applicationHandler) LoadConfiguration(w http.ResponseWriter, r *http.Request) {
	config, err := a.configService.LoadConfiguration(r.Context())
	if err != nil {
		_ = render.Render(w, r, newServiceErrResponse(err))
		return
	}

	_ = render.Render(w, r, newServerResponse("Configuration fetched successfully", config, http.StatusOK))
}

// CreateConfiguration
// @Summary Create a configuration
// @Description This endpoint creates a configuration
// @Tags Application
// @Accept  json
// @Produce  json
// @Param application body models.Configuration true "Configuration Details"
// @Success 200 {object} serverResponse{data=datastore.Configuration}
// @Failure 400,401,500 {object} serverResponse{data=Stub}
// @Security ApiKeyAuth
// @Router /configuration [post]
func (a *applicationHandler) CreateConfiguration(w http.ResponseWriter, r *http.Request) {
	var newConfig models.Configuration
	if err := util.ReadJSON(r, &newConfig); err != nil {
		_ = render.Render(w, r, newErrorResponse(err.Error(), http.StatusBadRequest))
		return
	}

	config, err := a.configService.CreateOrUpdateConfiguration(r.Context(), &newConfig)
	if err != nil {
		_ = render.Render(w, r, newServiceErrResponse(err))
		return
	}

	_ = render.Render(w, r, newServerResponse("Configuration created successfully", config, http.StatusOK))
}
