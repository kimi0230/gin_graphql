# Gin GraphQL


## Initiation
```shell
go mod init  
go mod tidy
```

## API Rules

## Build Server
```
make build
```

## GraphQL gqlgen
You could initialize a new project using the recommended folder structure by running this command

### Initiation
```shell
go run github.com/99designs/gqlgen init
```

`
├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go     - DO NOT EDIT
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go    - DO NOT EDIT
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit
`
### Finishing touches
At the top of our resolver.go, between package and import, add the following line:
``` go
//go:generate go run github.com/99designs/gqlgen
```
This magic comment tells go generate what command to run when we want to regenerate our code. 
To run go generate recursively over your entire project, use this command:`go generate ./...`


### Modify schema
1. modify your schema `graph/schema.graphqls`
2. `./script/gqlgen.sh`
3. modify `graph/resolver.go`

---
## Fix
1. `graph/prelude.resolvers.go:19:34: cannot refer to unexported name generated.__DirectiveResolver`
    * rollback the version of gqlparser from github.com/vektah/gqlparser/v2 v2.2.0 to github.com/vektah/gqlparser/v2 v2.1.0
    ```
    go mod edit -require github.com/vektah/gqlparser/v2@v2.1.0    
    go clean -i github.com/vektah/gqlparser/v2  
    go get github.com/vektah/gqlparser/v2@v2.1.0
    ```
    * https://github.com/99designs/gqlgen/issues/1402
    * https://stackoverflow.com/questions/24855081/how-do-i-import-a-specific-version-of-a-package-using-go-get
    * https://github.com/golang/go/issues/44129


## Reference
* [go.uber.org/ratelimit](https://pkg.go.dev/go.uber.org/ratelimit)
* [didip/tollbooth](https://github.com/didip/tollbooth)