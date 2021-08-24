#!/bin/bash
printf "\n Regenerating gqlgen files \n"
rm -f graph/generated/generated.go \
    graph/models/models_gen.go 
time go run -v github.com/99designs/gqlgen $1
printf "\n Regenerating gqlgenDone.\n\n"