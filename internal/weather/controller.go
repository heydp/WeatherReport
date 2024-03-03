package weather

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ExternalController struct {
	manager *Manager
}

func NewController(m *Manager) *ExternalController {
	return &ExternalController{
		manager: m,
	}
}

func (c *ExternalController) MountRoutes(router chi.Router) {
	router.With().Route("/report", func(r chi.Router) {
		r.MethodFunc(http.MethodPost, "/email", c.EmailReport)
	})
}

func (c *ExternalController) EmailReport(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Chi Server, emailing the report"))
}
