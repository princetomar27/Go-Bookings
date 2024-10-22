package handlers

import (
	"bookings/internal/config"
	"bookings/internal/models"
	"bookings/internal/render"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	//_, _ = fmt.Fprintf(w, "this is the home page")
	//render.RenderTemplate(w, "home.page.tmpl")
	render.RenderTemplateAdvanced(w, r, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some business logic here
	stringMap := make(map[string]string)
	stringMap["hello"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplateAdvanced(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateAdvanced(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Render the room page - generals
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateAdvanced(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Render the room page - majors
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateAdvanced(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateAdvanced(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start: %s, end: %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles requests for Availability and sends JSON Response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateAdvanced(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("y must be greater than 0")
		return 0, err
	}

	result := x / y
	return result, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("100 / 10 = %f", f))
}
