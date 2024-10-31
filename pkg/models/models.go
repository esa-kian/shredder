package models

type Field struct {
	Name         string
	DataType     string
	IsPrimaryKey bool
	IsRequired   bool
}

type Model struct {
	EntityName string
	Fields     []Field
}
