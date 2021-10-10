# godocmd

A simple program to write exported definitions of Go programs packages in Markdown Format.
It is using the `go/ast` to parse the relevant material and writing them in a concise format.

---

<details>
	<summary> <strong> Package main </strong> </summary>	

---

##### Functions:

1. [`GetMappedSyntaxTree`](./main.go#L28)
2. [`Scan`](./main.go#L162)

---
##### Structs

1. [`FuncDecl`](./type.go#L27)

	Methods:
	1. [`String`](./fmt.go#L9)
2. [`StructDecl`](./type.go#L20)

	Methods:
	1. [`String`](./fmt.go#L13)
3. [`FuncDecls`](./type.go#L35)

	Methods:
	1. [`String`](./fmt.go#L17)
4. [`StructDecls`](./type.go#L33)

	Methods:
	1. [`String`](./fmt.go#L25)
5. [`Package`](./type.go#L12)

	Methods:
	1. [`String`](./fmt.go#L42)
6. [`Packages`](./type.go#L5)

	Methods:
	1. [`String`](./fmt.go#L72)
7. [`Pos`](./type.go#L7)


---
</details>