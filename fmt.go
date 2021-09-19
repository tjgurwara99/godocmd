package main

import "fmt"

func (fd FuncDecl) String() string {
	return fd.Name
}

func (sd StructDecl) String() string {
	return sd.Name
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
	for index, sd := range sds {
		str += fmt.Sprintf("%d. %s\n", index+1, sd)
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
	if pkg.StructDecls != nil {
		str += `##### Structs

`
		str += fmt.Sprint(pkg.StructDecls)
		str += "\n---\n"
	}
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
