package generator

type fieldType string

const fieldTypeObjectID fieldType = "objectID"
const fieldTypeString fieldType = "string"
const fieldTypeDate fieldType = "date"

type schema struct {
	Models []modelSchema `yaml:"models"`
}

type modelSchema struct {
	Name   string        `yaml:"name"`
	Fields []fieldSchema `yaml:"fields"`
}

type fieldSchema struct {
	Name string    `yaml:"name"`
	Type fieldType `yaml:"type"`
}
