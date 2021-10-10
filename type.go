package main

import "go/token"

type Packages map[string]Package

type Pos struct {
	FileName string
	Line     int
}

type Package struct {
	Name        string
	Parent      string
	Description string
	StructDecls StructDecls
	FuncDecls   FuncDecls
}

type StructDecl struct {
	Name        string
	Pos         Pos
	Fset        *token.FileSet
	FuncDecls   FuncDecls
	Description string
}

type FuncDecl struct {
	Name        string
	Pos         Pos
	Fset        *token.FileSet
	Description string
}

type StructDecls map[string]StructDecl

type FuncDecls []FuncDecl
