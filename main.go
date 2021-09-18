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

func GetPackages(currentPath string, fset *token.FileSet) map[string]*ast.Package {

	pkgs, err := parser.ParseDir(fset, currentPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pkgs
}

func MakeTreeToPrint(pkgs map[string]*ast.Package) map[string]interface{} {
	dictionary := make(map[string]interface{})
	for name, synTree := range pkgs {
		fmt.Printf("%v\n", name)
		ast.PackageExports(synTree)
		var functionList []string
		ast.Inspect(synTree, func(n ast.Node) bool {
			var s string
			switch x := n.(type) {
			case *ast.FuncDecl:
				s = x.Name.Name
				if strings.HasPrefix(s, "Test") {
					return true
				}
				functionList = append(functionList, s)
			}
			return true
		})
		dictionary[name] = functionList
	}
	return dictionary
}

func main() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	pkgs := GetPackages(currentPath, fset)
	fmt.Print(MakeTreeToPrint(pkgs))

	// ast.Print(fset, pkgs["main"])
}
