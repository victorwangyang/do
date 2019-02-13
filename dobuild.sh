#!/bin/bash

echo "starting build do project ....."

go build -o ./do/do ./do/do.go
echo "do building.......OK"
go build -o ./master/master ./master/master.go
echo "master building.......OK"
go build -o ./node/node ./node/node.go
echo "node building.......OK"

echo "ending......."