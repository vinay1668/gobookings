package handlers

import (
	"net/http"

	"github.com/vinay1668/gobookings/pkg/config"
	"github.com/vinay1668/gobookings/pkg/models"
	"github.com/vinay1668/gobookings/pkg/render"
)

//TemplateData holds data sent from handlers to templates

//Repo the repository used by the handlers
var Repo *Repository


//Repository is the repository type
type Repository struct {
	App *config.AppConfig

}
//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
      return &Repository{
         App:a,
	  }
}
// NewHandlers sets the repository for the handler
func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
    render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{});
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	//perform some bussiness logic

	//
	stringMap := make(map[string]string)
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")


	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	});

}


