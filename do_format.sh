#!/bin/sh

for i in "events tour gambler results"
do
    cd ${i}
    go fmt -n
    cd ..
done
