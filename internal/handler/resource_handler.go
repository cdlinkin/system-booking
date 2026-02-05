package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cdlinkin/system-booking/internal/service"
)

type ResourceHandler struct {
	resourceService service.ResourceService
}

func NewResourceHandler(resourceService service.ResourceService) *ResourceHandler {
	return &ResourceHandler{resourceService: resourceService}
}

func (h *ResourceHandler) GetResources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	available := r.URL.Query().Get("available")
	onlyAvailable := available == "true"

	resources, err := h.resourceService.GetResources(r.Context(), onlyAvailable)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resources)
	w.WriteHeader(http.StatusOK)
}
