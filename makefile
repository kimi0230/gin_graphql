all: build

bin_file=gin_graphql

build: clean
	./script/gobuild.sh

clean:
	rm -rf build/${bin_file}

build_graphql:
	./script/gqlgen.sh

clean_build_graphql:
	rm -rf graph/generated/generated.go
	rm -rf graph/models/models_gen.go
	rm -rf graph/schema.resolvers.go

.PHONY: clean build all