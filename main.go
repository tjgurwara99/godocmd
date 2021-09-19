package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getAstPackage(sourcePath string, fset *token.FileSet) map[string]*ast.Package {
	pkgs, err := parser.ParseDir(fset, sourcePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pkgs
}

func MakeTreeToPrint(pkgs map[string]*ast.Package, fset *token.FileSet) Packages {
	packages := make(map[string]Package)
	for name, syntaxTree := range pkgs {
		ast.PackageExports(syntaxTree)
		var functionList FuncDecls
		var structList StructDecls
		ast.Inspect(syntaxTree, func(n ast.Node) bool {
			var fd FuncDecl
			switch x := n.(type) {
			case *ast.FuncDecl:
				if x.Recv != nil {
					return true
				}
				fd.Name = x.Name.Name
				if strings.HasPrefix(fd.Name, "Test") || strings.HasPrefix(fd.Name, "Example") {
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

		packages[name] = Package{
			Name:        name,
			FuncDecls:   functionList,
			StructDecls: structList,
		}
	}
	return packages
}

func main() {
	flag.Parse()

	if len(flag.Args()) > 1 {
		w := flag.CommandLine.Output()

		fmt.Fprintf(w, `Error: %s only accepts one argument.
If no argument is provided it looks at the current directory by default.

`, os.Args[0])
		flag.Usage()
		os.Exit(1)
	}

	sourcePath := flag.Arg(0)

	fmt.Print(sourcePath)

	if sourcePath == "" {
		var err error
		sourcePath, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	scanned := Scan(sourcePath)
	fmt.Print(scanned)

}

func scanDir(path string) Packages {
	fset := token.NewFileSet()
	astPkgs := getAstPackage(path, fset)
	return MakeTreeToPrint(astPkgs, fset)
}

func Scan(sourcePath string) Packages {
	if !recursive {
		return scanDir(sourcePath)
	}
	packages := scanDir(sourcePath)
	err := filepath.Walk(sourcePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			subpackages := scanDir(path)
			for key, value := range subpackages {
				packages[key] = value
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return packages
}
