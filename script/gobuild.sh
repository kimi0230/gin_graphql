#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
outputPath="build"
app="gin_graphql"
output="$outputPath/$app"
src="$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"
