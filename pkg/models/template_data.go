package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloateMap map[string]float32
	Date      map[string]any
	CSRFToken string
	Flash     string
	Error     string
}
