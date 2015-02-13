#!/bin/sh

for i in events tour gambler results
do
    echo "$i"
    cd ${i}
    go fmt -n
    cd ..
done
