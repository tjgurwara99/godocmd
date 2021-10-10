package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/mod/modfile"
)

func getAstPackage(sourcePath string, fset *token.FileSet) map[string]*ast.Package {
	pkgs, err := parser.ParseDir(fset, sourcePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	return pkgs
}

func GetMappedSyntaxTree(pkgs map[string]*ast.Package, fset *token.FileSet) Packages {
	packages := make(Packages)
	for name, pkgSyntaxTree := range pkgs {
		ast.PackageExports(pkgSyntaxTree)
		var functionList FuncDecls
		structList := make(StructDecls)
		var description string
		for _, file := range pkgSyntaxTree.Files {
			if file.Doc == nil {
				continue
			}
			for _, doc := range file.Doc.List {
				text := doc.Text
				reg := regexp.MustCompile(`^\/\/ |^\/\* |^\/\/| \*\/$|\*\/$`)
				description += reg.ReplaceAllString(text, " ")
			}
		}
		ast.Inspect(pkgSyntaxTree, func(n ast.Node) bool {
			var fd FuncDecl
			switch x := n.(type) {
			case *ast.FuncDecl:
				fd.Name = x.Name.Name
				fd.Pos = Pos{
					Line:     fset.Position(x.Pos()).Line,
					FileName: fset.Position(x.Pos()).Filename,
				}
				if x.Recv != nil {
					if s, ok := x.Recv.List[0].Type.(*ast.Ident); ok {
						if val, ok2 := structList[s.Name]; ok2 {
							val.FuncDecls = append(structList[s.Name].FuncDecls, fd)
							return true
						}
						structList[s.Name] = StructDecl{
							Name:      s.Name,
							FuncDecls: FuncDecls{fd},
						}
					}
					return true
				}
				// ignore Test, Example and Benchmark prefixed functions
				if strings.HasPrefix(fd.Name, "Test") || strings.HasPrefix(fd.Name, "Example") || strings.HasPrefix(fd.Name, "Benchmark") {
					return true
				}
				functionList = append(functionList, fd)
			case *ast.TypeSpec:
				var sd StructDecl
				sd.Name = x.Name.Name
				sd.Pos = Pos{
					Line:     fset.Position(x.Pos()).Line,
					FileName: fset.Position(x.Pos()).Filename,
				}
				if _, ok := x.Type.(*ast.StructType); ok {
					if x, ok := structList[sd.Name]; ok {
						x.Pos = sd.Pos
						structList[sd.Name] = x
						return true
					}
					structList[sd.Name] = sd
				}
			}
			return true
		})

		packages[name] = Package{
			Name:        name,
			Description: description,
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
If no argument is provided godocmd will look at the current directory by default.

`, os.Args[0])
		flag.Usage()
		os.Exit(1)
	}

	sourcePath = flag.Arg(0)

	if sourcePath == "" {
		var err error
		sourcePath, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	if module == "" {
		modfileData, err := ioutil.ReadFile(sourcePath + "/go.mod")
		if err != nil {
			log.Print("Unable to read go.mod file in the source directory provided. Please use -module flag to specify module path, example godocmd -module github.com/tjgurwara99/godocmd")
			flag.Usage()
			log.Fatal()
		}
		module = modfile.ModulePath(modfileData)
	}

	scanned := Scan(sourcePath)
	fmt.Print(scanned)

}

func scanDir(path string) Packages {
	fset := token.NewFileSet()
	astPkgs := getAstPackage(path, fset)
	return GetMappedSyntaxTree(astPkgs, fset)
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
