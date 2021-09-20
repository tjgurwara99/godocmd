package main

import "fmt"

func (fd FuncDecl) String() string {
	return fmt.Sprintf("`%s`", fd.Name)
}

func (sd StructDecl) String() string {
	return fmt.Sprintf("`%s`", sd.Name)
}

func (fds FuncDecls) String() string {
	str := ""
	for index, fd := range fds {
		str += fmt.Sprintf("%d. %s\n", index+1, fd)
	}
	return str
}

func (sds StructDecls) String() string {
	str := ""
	index := 1
	for _, sd := range sds {
		str += fmt.Sprintf("%d. %s\n", index, sd)

		if sd.FuncDecls != nil {
			str += "\tMethods:\n"
		}
		for i, f := range sd.FuncDecls {
			str += fmt.Sprintf("\t%d. %s\n", i+1, f)
		}
		index++
	}
	return str
}

func (pkg Package) String() string {
	str := fmt.Sprintf(`<details>
	<summary> <strong> Package %s </strong> </summary>	

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
		str += `##### Structs

`
		str += fmt.Sprint(pkg.StructDecls)
		str += "\n---\n"
	}
	str += "</details>"
	return str
}

func (pkgs Packages) String() string {
	var str string
	for _, pkg := range pkgs {
		if len(pkg.FuncDecls) != 0 || len(pkg.StructDecls) != 0 {
			str += fmt.Sprint(pkg)
		}
	}
	return str
}
