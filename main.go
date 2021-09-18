package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

type FuncDecl struct {
	Name string
	Pos  string
}

func GetPackages(currentPath string, fset *token.FileSet) map[string]*ast.Package {
	pkgs, err := parser.ParseDir(fset, currentPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pkgs
}

func MakeTreeToPrint(pkgs map[string]*ast.Package, fset *token.FileSet) map[string]interface{} {
	dictionary := make(map[string]interface{})
	for name, syntaxTree := range pkgs {
		fmt.Printf("%v\n", name)
		ast.PackageExports(syntaxTree)
		var functionList []FuncDecl
		ast.Inspect(syntaxTree, func(n ast.Node) bool {
			var fd FuncDecl
			switch x := n.(type) {
			case *ast.FuncDecl:
				fd.Name = x.Name.Name
				if strings.HasPrefix(fd.Name, "Test") {
					return true
				}
				functionList = append(functionList, fd)
			}
			return true
		})

		dictionary[name] = functionList
	}
	// pos := fset.Position(150)
	return dictionary
}

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	pkgs := GetPackages(currentPath, fset)
	fmt.Print(MakeTreeToPrint(pkgs, fset))
}
