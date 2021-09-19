package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func getAstPackage(sourcePath string, fset *token.FileSet) map[string]*ast.Package {
	pkgs, err := parser.ParseDir(fset, sourcePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pkgs
}

func MakeTreeToPrint(pkgs map[string]*ast.Package, fset *token.FileSet) map[string]Package {
	dictionary := make(map[string]Package)
	for name, syntaxTree := range pkgs {
		fmt.Printf("Package %v:\n\tDescription: \n", name)
		ast.PackageExports(syntaxTree)
		var functionList FuncDecls
		var structList StructDecls
		ast.Inspect(syntaxTree, func(n ast.Node) bool {
			var fd FuncDecl
			// var sd StructDecl
			switch x := n.(type) {
			case *ast.FuncDecl:
				if x.Recv != nil {
					return true
				}
				fd.Name = x.Name.Name
				if strings.HasPrefix(fd.Name, "Test") {
					return true
				}
				functionList = append(functionList, fd)
			case *ast.TypeSpec:
				var sd StructDecl
				if _, ok := x.Type.(*ast.StructType); ok {
					sd.Name = x.Name.Name
				}
				structList = append(structList, sd)
			}
			return true
		})

		dictionary[name] = Package{
			FuncDecls:   functionList,
			StructDecls: structList,
		}
	}
	return dictionary
}

func main() {

	flag.BoolVar(&recursive, "r", false, "Recursively traverse the source")
	flag.StringVar(&sourcePath, "path", "", "Path to the source traversal, defaults to current directory")

	flag.Parse()

	if sourcePath == "" {
		var err error
		sourcePath, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	fset := token.NewFileSet()
	pkgs := getAstPackage(sourcePath, fset)
	data := MakeTreeToPrint(pkgs, fset)

	for _, pkg := range data {
		fmt.Print("\t\tFunctions:\n")
		fmt.Print(pkg.FuncDecls)
		fmt.Print("\t\tStructs:\n")
		fmt.Print(pkg.StructDecls)
	}
}
