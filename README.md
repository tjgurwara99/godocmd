# godocmd

A simple program to write exported definitions of Go programs packages in Markdown Format.
It is using the `go/ast` to parse the relevant material and writing them in a concise format.

---
<details>
	<summary> <strong> Package main </strong> </summary>	

---

##### Functions:

1. [`GetMappedSyntaxTree`](./main.go#L28)
2. [`Scan`](./main.go#L146)

---
##### Structs

1. [`Pos`](./type.go#L7)
2. [`FuncDecl`](./type.go#L27)
	Methods:
	1. [`String`](./fmt.go#L9)
3. [`StructDecl`](./type.go#L20)
	Methods:
	1. [`String`](./fmt.go#L13)
4. [`FuncDecls`](#L0)
	Methods:
	1. [`String`](./fmt.go#L17)
5. [`StructDecls`](#L0)
	Methods:
	1. [`String`](./fmt.go#L25)
6. [`Package`](./type.go#L12)
	Methods:
	1. [`String`](./fmt.go#L42)
7. [`Packages`](#L0)
	Methods:
	1. [`String`](./fmt.go#L72)

---
</details>