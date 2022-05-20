package handlers

import (
	"net/http"

	"github.com/calvarado2004/bookings/pkg/config"
	"github.com/calvarado2004/bookings/pkg/models"
	"github.com/calvarado2004/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//	About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//Perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Specials(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "specials.page.html", &models.TemplateData{})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.html", &models.TemplateData{})
}
