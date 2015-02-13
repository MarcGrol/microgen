#!/bin/sh

go fmt
for i in events tour gambler results
do
    cd ${i}
    go fmt 
    cd ..
done
