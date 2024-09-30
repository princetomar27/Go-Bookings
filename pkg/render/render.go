package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

var templateCacheMap = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, searchTemp string) {

	var temp *template.Template
	var err error

	// Check if the template exists in the map or not
	_, isTempExists := templateCacheMap[searchTemp]

	if !isTempExists {
		log.Println("Creating template and adding in Cache")
		// if it exists in the map
		err = createTemplateCache(searchTemp)
		if err != nil {
			log.Println("Error creating template cache", err)
		}
	} else {
		// if it doesn't exist in the map
		log.Printf("Template %s does not exist in cache", searchTemp)
		log.Println("Template exist in Cache")
	}

	temp = templateCacheMap[searchTemp]
	err = temp.Execute(w, nil)
	if err != nil {
		log.Println("Error creating template cache", err)
	}
}

func createTemplateCache(t string) error {

	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the templates
	tmpl, err := template.ParseFiles(templates...) // Iterating over templates array/slice,
	// and passing each entry to template.ParseFiles() as individual entry

	if err != nil {
		return err
	}

	// Else add the template to the template Cache map
	templateCacheMap[t] = tmpl

	return nil
}
