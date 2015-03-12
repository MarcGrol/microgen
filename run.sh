#!/bin/sh

find . -name "*.dat" -exec rm -f {} \;
find . -name "*.db" -exec rm -f {} \;

pkill microgen

microgen -service=proxy -port=8080 &
microgen -service=tour -port=8081 &
microgen -service=gambler -port=8082 &
microgen -service=news -port=8083 &
microgen -service=collector -port=8084 &

ps -eaf|grep microgen |grep -v grep
