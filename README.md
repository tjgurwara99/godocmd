# godocmd

A simple program to write exported definitions of Go programs packages in Markdown Format.
It is using the `go/ast` to parse the relevant material and writing them in a concise format.

---

<details>
	<summary> <strong> Package main </strong> </summary>	

---

##### Functions:

1. [`GetMappedSyntaxTree`](./main.go#L28)
2. [`Scan`](./main.go#L148)

---
##### Structs

1. [`Pos`](./type.go#L7)
2. [`FuncDecl`](./fmt.go#L8)
	Methods:
	1. [`String`](./fmt.go#L8)
3. [`StructDecl`](./fmt.go#L12)
	Methods:
	1. [`String`](./fmt.go#L12)
4. [`FuncDecls`](./fmt.go#L16)
	Methods:
	1. [`String`](./fmt.go#L16)
5. [`StructDecls`](./fmt.go#L24)
	Methods:
	1. [`String`](./fmt.go#L24)
6. [`Package`](./fmt.go#L41)
	Methods:
	1. [`String`](./fmt.go#L41)
7. [`Packages`](./fmt.go#L71)
	Methods:
	1. [`String`](./fmt.go#L71)

---
</details>