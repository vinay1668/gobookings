package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/vinay1668/gobookings/pkg/config"
	"github.com/vinay1668/gobookings/pkg/models"
)

var app *config.AppConfig


func NewTemplates(a *config.AppConfig){
   app = a;
}

func addDefaultData(td *models.TemplateData) *models.TemplateData{
	return td;
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData){
     var tc map[string]*template.Template
	if app.UseCache{
		tc = app.TemplateCache;
	 }else{
		tc, _ = CreateTemplateCache()
	 }
// get the template cache from app config



	//get requested template from cache

    t, ok := tc[tmpl];
	if !ok{
       log.Fatal("Could not get template from Template Cache!")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)

	err := t.Execute(buf, td)

	if err != nil{
		log.Println(err)
	}
	
	//render the template

	_, err = buf.WriteTo(w)
	if err != nil{
		log.Println(err)
	}
	// parsedTemplate , _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl");
	// err := parsedTemplate.Execute(w, nil);
	// if err != nil {
    //       fmt.Println("Error parsing template:", err);
	// 	  return;
	// }
}


func CreateTemplateCache() (map[string]*template.Template, error){
	myCache := map[string]*template.Template{}

	// get all of the files names *.page.tmpl from ./templates

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return myCache,err
	}

	//range through all the files ending with *.page.tmpl

	for _, page:= range pages {
		 name:= filepath.Base(page)
		 ts, err := template.New(name).ParseFiles(page)
		 if err != nil{
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil{
			return myCache,err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return myCache,err
			}
		}

		myCache[name] = ts
		 
	}
	return myCache, nil
}