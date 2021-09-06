# Gin GraphQL
A GraphQL server with Golang, gorm, gqlgen.

## Build Server
Copy `.env.example` to `.env` and `.env.dev`

```shell

# create tables
make migrate

# create GraphQL file
make build_graphql 

# build server
make build

# start server
.build/gin_graphql
```

---
## Notes
### GraphQL 
#### [Schema][6]
* https://graphql.org/learn/schema/

#### [99designs/gqlgen][1]
https://gqlgen.com/
You could initialize a new project using the recommended folder structure by running this command

#### Initiation
```shell
go run github.com/99designs/gqlgen init
```

```
├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go     - DO NOT EDIT !
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go    - DO NOT EDIT !
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit
```

#### Finishing touches
At the top of our resolver.go, between package and import, add the following line:
``` go
//go:generate go run github.com/99designs/gqlgen
```
This magic comment tells go generate what command to run when we want to regenerate our code. 
To run go generate recursively over your entire project, use this command:`go generate ./...`


#### Modify schema
1. modify your schema `graph/schema.graphqls`
2. run gqlgen `./script/gqlgen.sh` or `go run -v github.com/99designs/gqlgen`
3. modify resolvers `graph/resolver.go`

---

### Database (Mysql)
#### Migrate
`make migrate mode={auto | drop | refresh}` or `go run  database/migrate.go -m=auto`

---
### Go commands
```shell
# Initiation Go project
go mod init  
go mod tidy

# download Go package
go mod download
```
---
### Fix
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

---

## Reference
1. [99designs/gqlgen][1]
2. [go.uber.org/ratelimit][2]
3. [didip/tollbooth][3]
4. [EQuimper/youtube-golang-graphql-tutorial][4]
5. [wtlin1228/unasees][5]
6. [schema][6]
7. [gqlgen gin][7]
8. [blog.laisky.com][8]

[1]: https://github.com/99designs/gqlgen 
"99designs/gqlgen"
[2]: https://pkg.go.dev/go.uber.org/ratelimit
"go.uber.org/ratelimit"
[3]: https://github.com/didip/tollbooth
"didip/tollbooth"
[4]: https://github.com/EQuimper/youtube-golang-graphql-tutorial
"EQuimper/youtube-golang-graphql-tutorial"
[5]: https://github.com/wtlin1228/unasees
"wtlin1228/unasees"
[6]: https://graphql.org/learn/schema/
"schema"
[7]: https://gqlgen.com/recipes/gin/
"gqlgen gin"
[8]: https://blog.laisky.com/p/gqlgen/#%E5%AE%9A%E4%B9%89+schema-Hfxfd
