#!/bin/sh

bin=$(pwd | rev | cut -d '/' -f 1 | rev)
go build -o cmd/$bin cmd/main.go

echo "compiled: cmd/${bin}"
