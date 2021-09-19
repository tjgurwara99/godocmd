package main

import "fmt"

func (fd FuncDecl) String() string {
	return fmt.Sprintf("<li> %s </li>\n", fd.Name)
}

func (sd StructDecl) String() string {
	return fmt.Sprintf("<li> %s </li>\n", sd.Name)
}

func (fds FuncDecls) String() string {
	str := "<ol>\n"
	for _, fd := range fds {
		str += fmt.Sprint(fd)
	}
	str += "</ol>"
	return str
}

func (sds StructDecls) String() string {
	str := "<ol>\n"
	for _, sd := range sds {
		str += fmt.Sprint(sd)
	}
	str += "</ol>"
	return str
}

func (pkg Package) String() string {
	str := fmt.Sprintf(`<details>
	<summary> <strong> Package %s </strong> </summary>	
`, pkg.Name)
	if pkg.Description != "" {
		str += fmt.Sprintf(`<p> 
			<details> <summary> Description </summary>
			%s
			</details>
			</p>
`, pkg.Description)
	}
	if pkg.FuncDecls != nil {
		str += `<p> 
		<details> <summary> Functions </summary>
		`
		str += fmt.Sprint(pkg.FuncDecls)
		str += `</details>
		</p>`
	}
	if pkg.StructDecls != nil {
		str += `<p> 
		<details> <summary> Structs  </summary>`
		str += fmt.Sprint(pkg.StructDecls)
		str += `</details>
		</p>`
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
