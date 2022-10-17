package main

import (
	"fmt"
	"go/doc"
	"go/token"
	"strings"

	"github.com/tjgurwara99/compose"
	"golang.org/x/tools/go/packages"
)

type Node struct {
	Name        string
	Package     *packages.Package
	Doc         *doc.Package
	SubPackages Tree
}

type Tree []*Node

func NewDocTree(pkgs []*packages.Package) Tree {
	var pathTree Tree
	pkgPaths := compose.Map(pkgs, func(pkg *packages.Package) string { return pkg.ID })
	for _, pkg := range pkgPaths {
		paths := strings.Split(strings.TrimPrefix(pkg, pkgs[0].ID), "/")
		paths[0] = pkgs[0].ID + paths[0]
		pathTree = AppendToTree(pathTree, paths)
	}
	TraverseAndInject(pathTree, "", pkgs)
	return pathTree
}

func TraverseAndInject(root Tree, before string, pkgs []*packages.Package) Tree {
	if root == nil {
		return root
	}
	for _, node := range root {
		var path string
		if before == "" {
			path = pkgs[0].ID
		} else {
			path = before + "/" + node.Name
		}
		TraverseAndInject(node.SubPackages, path, pkgs)
		pkg, ok := compose.Find(pkgs, func(pkg *packages.Package) bool {
			return fmt.Sprintf("%s/%s", before, node.Name) == pkg.ID
		})
		if !ok {
			if before == "" {
				node.Package = pkgs[0] // root package
				node.Doc, _ = findInFile(pkgs[0], fset)
			}
			continue
		}
		node.Package = pkg
		node.Doc, _ = findInFile(pkg, fset)
	}
	return root
}

func AppendToTree(root Tree, pkgsIDs []string) Tree {
	if len(pkgsIDs) == 0 {
		return root
	}
	var i int
	for i = 0; i < len(root); i++ {
		if root[i].Name == pkgsIDs[0] {
			break
		}
	}
	if i == len(root) {
		root = append(root, &Node{Name: pkgsIDs[0]})
	}
	root[i].SubPackages = AppendToTree(root[i].SubPackages, pkgsIDs[1:])
	return root
}

func findInFile(pkg *packages.Package, fset *token.FileSet) (*doc.Package, error) {
	p, err := doc.NewFromFiles(fset, pkg.Syntax, pkg.Module.Path)
	if err != nil {
		return nil, err
	}
	return p, nil
}
