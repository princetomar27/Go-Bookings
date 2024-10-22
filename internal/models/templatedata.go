package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // when we're not sure what might be the data type can be - a struct ex.
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
