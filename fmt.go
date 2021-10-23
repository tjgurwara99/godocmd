package main

import (
	"fmt"
	"sort"
	"strings"
)

func (fd FuncDecl) String() string {
	return fmt.Sprintf("`%s`", fd.Name)
}

func (sd StructDecl) String() string {
	return fmt.Sprintf("`%s`", sd.Name)
}

func (fds FuncDecls) String() string {
	str := ""
	for index, fd := range fds {
		str += fmt.Sprintf("%d. [%s](%s): %s\n", index+1, fd, fd.Pos.String(), fd.Description)
	}
	return str
}

func (sds StructDecls) String() string {
	str := ""
	index := 1
	for _, sd := range sds {
		str += fmt.Sprintf("%d. [%s](%s): %s\n\n", index, sd, sd.Pos.String(), sd.Description)

		if sd.FuncDecls != nil {
			str += "\tMethods:\n"
		}
		for i, f := range sd.FuncDecls {
			str += fmt.Sprintf("\t%d. [%s](%s): %s\n", i+1, f, f.Pos.String(), f.Description)
		}
		index++
	}
	return str
}

func (pkg Package) String() string {
	str := fmt.Sprintf(`<details>
	<summary> <strong> %s </strong> </summary>	

---

`, pkg.Name)
	if pkg.Description != "" {
		str += fmt.Sprintf(`##### Description: %s
`, pkg.Description)
		str += "\n---\n"
	}
	if pkg.FuncDecls != nil {
		str += `##### Functions:

`
		str += fmt.Sprint(pkg.FuncDecls)
		str += "\n---\n"
	}
	if len(pkg.StructDecls) != 0 {
		str += `##### Types

`
		str += fmt.Sprint(pkg.StructDecls)
		str += "\n---\n"
	}
	str += "</details>"
	return str
}

func (pkgs Packages) String() string {
	var str string = "# Packages:\n\n"
	var pkgSlice []Package
	for _, pkg := range pkgs {
		pkgSlice = append(pkgSlice, pkg)
	}
	sort.Slice(pkgSlice, func(i, j int) bool { return pkgSlice[i].Name < pkgSlice[j].Name })
	for _, pkg := range pkgSlice {
		if len(pkg.FuncDecls) != 0 || len(pkg.StructDecls) != 0 {
			str += fmt.Sprint(pkg)
		}
	}
	return str
}

func (p *Pos) String() string {
	str := fmt.Sprintf("%s#L%d", strings.Replace(p.FileName, sourcePath, ".", 1), p.Line)
	return str
}
