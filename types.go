package main

type Package struct {
	FileName   []string
	Package    string
	Structures []*Struct
	Funcs      []*StructFunc
}

type Struct struct {
	Name   string
	Fields []interface{}
}

type Field struct {
	Name          string
	Pointer       bool
	Slice         bool
	Interface     bool
	KindName      string
	KindPkg       string
	PrimitiveType string
	Generic       *Field
}

type StructFunc struct {
	Struct string
	Name   string
}
