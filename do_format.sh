#!/bin/sh

go fmt
for i in spec store bus tourApp/events tourApp/tour tourApp/gambler tourApp/results
do
    cd ${i}
    go fmt 
    cd -
done
