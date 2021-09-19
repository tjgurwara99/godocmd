package main

import "fmt"

type Package struct {
	Name        string
	Parent      string
	StructDecls StructDecls
	FuncDecls   FuncDecls
}

type StructDecl struct {
	Name string
	Pos  string
}

type FuncDecl struct {
	Name string
	Pos  string
}

type StructDecls []StructDecl

type FuncDecls []FuncDecl

func (fd FuncDecl) String() string {
	return fmt.Sprintf("\t\t\t%s\n", fd.Name)
}

func (sd StructDecl) String() string {
	return fmt.Sprintf("\t\t\t%s\n", sd.Name)
}

func (fds FuncDecls) String() string {
	var str string
	for _, fd := range fds {
		str += fmt.Sprint(fd)
	}
	return str
}

func (sds StructDecls) String() string {
	var str string
	for _, sd := range sds {
		str += fmt.Sprint(sd)
	}
	return str
}

func (pkg Package) String() string {
	var str string
	if pkg.FuncDecls != nil {
		str = "\t\tFunctions:\n"
		str += fmt.Sprint(pkg.FuncDecls)
	}
	if pkg.StructDecls != nil {
		str += "\t\tStructs:\n"
		str += fmt.Sprint(pkg.StructDecls)
	}
	return str
}
