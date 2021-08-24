# Gin GraphQL


## Initiation
```shell
go mod init  
go mod tidy
```

## API Rules

## GraphQL gqlgen
You could initialize a new project using the recommended folder structure by running this command

```shell
go run github.com/99designs/gqlgen init
```

## fix
1. 
`graph/prelude.resolvers.go:19:34: cannot refer to unexported name generated.__DirectiveResolver`
* https://github.com/99designs/gqlgen/issues/1402
* rollback the version of gqlparser from github.com/vektah/gqlparser/v2 v2.2.0 to github.com/vektah/gqlparser/v2 v2.1.0
```
go mod edit -require github.com/vektah/gqlparser/v2@v2.1.0    
go clean -i github.com/vektah/gqlparser/v2  
go get github.com/vektah/gqlparser/v2@v2.1.0
```


## Reference
* [go.uber.org/ratelimit](https://pkg.go.dev/go.uber.org/ratelimit)
* [didip/tollbooth](https://github.com/didip/tollbooth)