## Package:

- A module is a collection of go packages.
- A package is a directory of .go files. Using packages, you organize your code into reusable units.
- We can add a module to go project or upgrade the module version.
- [Go Module vs Package](https://stackoverflow.com/questions/61940117/go-modules-vs-package)

```
MyModules
├── Module1
|  ├── LICENSE
|  ├── go.mod
|  └── go.sum
|  └── package1
|       └── func1.go
|       └── func2.go
|
├── Module2
   ├── LICENSE
   ├── go.mod
   └── go.sum
   └── package1
        └── func1.go
        └── func2.go
```

- A package in Go is essentially a named collection of one or more related .go files. 
- In Go, the primary purpose of packages is to help you isolate and reuse code.

- Every .go file that you write should begin with a package {name} statement which indicates the name of the package that the file is a part of. 
