# godocmd

A simple program to write exported definitions of Go programs packages in Markdown Format.
It is using the `go/ast` to parse the relevant material and writing them in a concise format.

---
<details>
	<summary> <strong> Package main </strong> </summary>	

---

##### Functions:

1. [`GetMappedSyntaxTree`](./main.go#L28)
2. [`Scan`](./main.go#L154)

---
##### Structs

1. [`Packages`](./type.go#L5)

	Methods:
	1. [`String`](./fmt.go#L72)
2. [`Pos`](./type.go#L7)

3. [`FuncDecl`](./type.go#L27)

	Methods:
	1. [`String`](./fmt.go#L9)
4. [`StructDecl`](./type.go#L20)

	Methods:
	1. [`String`](./fmt.go#L13)
5. [`FuncDecls`](#L0)

	Methods:
	1. [`String`](./fmt.go#L17)
6. [`StructDecls`](./type.go#L33)

	Methods:
	1. [`String`](./fmt.go#L25)
7. [`Package`](./type.go#L12)

	Methods:
	1. [`String`](./fmt.go#L42)

---
</details>