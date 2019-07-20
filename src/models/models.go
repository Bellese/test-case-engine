package models

// ConfigField contains all information needed to define one piece of data that we are generating
type ConfigField struct {
	Name string
	Type string
	Min  int
	Max  int
}

// Config contains all information that is needed to generate test data
type Config struct {
	Title  string
	Output string
	Total  int
	Fields []ConfigField
}
