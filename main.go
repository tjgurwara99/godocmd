package main

import (
	"flag"
	"fmt"
	"go/doc/comment"
	"go/token"
	"io"
	"os"

	"golang.org/x/tools/go/packages"
)

var fset = token.NewFileSet()

func main() {
	const mode packages.LoadMode = packages.NeedModule |
		packages.NeedName |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo

	flag.Parse()
	cfg := &packages.Config{Fset: fset, Mode: mode}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	pathTree := NewDocTree(pkgs)
	FPrint(os.Stdout, pathTree, "", "")
}

func FPrint(w io.Writer, tree Tree, indent, prefix string) {
	if prefix == "" {
		prefix = "##"
	}
	for _, pkg := range tree {
		printPkgDoc(pkg, prefix, w)
		fmt.Fprintln(w, "\nSubpackages:")
		newWriter := NewWriter(w, indent+"\t")
		FPrint(newWriter, pkg.SubPackages, indent+"\t", prefix+"#")
	}
}

func printPkgDoc(pkg *Node, prefix string, w io.Writer) {
	parser := pkg.Doc.Parser()
	var pkgComment string
	if pkg.Doc != nil {
		pkgComment = fmt.Sprintf("%s: %s", pkg.Doc.Name, pkg.Doc.Doc)
		if pkg.Doc.Doc == "" {
			pkgComment = fmt.Sprintf("%s: No package documentation provided.", pkg.Doc.Name)
		}
	} else {
		pkgComment = fmt.Sprintf("%s: No package documentation provided.", pkg.Name)
	}
	doc := parser.Parse(pkgComment)
	printer := comment.Printer{}
	printer.TextPrefix = prefix + " "
	printer.TextCodePrefix = "```"
	something := printer.Text(doc)
	w.Write([]byte(something))

	if pkg.Doc == nil {
		return
	}

	if len(pkg.Doc.Funcs) > 0 {
		w.Write(printer.Text(parser.Parse("Functions:")))
		printer.TextPrefix += " "
		for _, fn := range pkg.Doc.Funcs {
			var fnCmt string
			if fn.Doc == "" {
				fn.Doc = "No function documentation provided."
			}
			fnCmt = fmt.Sprintf("%s: %s", fn.Name, fn.Doc)
			fnDoc := parser.Parse(fnCmt)
			fDoc := printer.Markdown(fnDoc)
			w.Write([]byte(fDoc))
		}
	}
	if len(pkg.Doc.Consts) > 0 {
		printer.TextPrefix = printer.TextPrefix[:len(printer.TextPrefix)-1]
		w.Write([]byte(printer.Text(parser.Parse("Types:"))))
		printer.TextPrefix += " "
		for _, t := range pkg.Doc.Types {
			var tCmt string
			if t.Doc == "" {
				t.Doc = "No type documentation provided."
			}
			tCmt = fmt.Sprintf("%s: %s", t.Name, t.Doc)
			tDoc := parser.Parse(tCmt)
			w.Write(printer.Text(tDoc))
		}
	}
}
