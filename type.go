package main

type Packages map[string]Package

type Package struct {
	Name        string
	Parent      string
	Description string
	StructDecls StructDecls
	FuncDecls   FuncDecls
}

type StructDecl struct {
	Name      string
	Pos       string
	FuncDecls FuncDecls
}

type FuncDecl struct {
	Name string
	Pos  string
}

type StructDecls map[string]StructDecl

type FuncDecls []FuncDecl
