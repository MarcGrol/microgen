#!/bin/sh

MICROGEN=/home/mgrol/go/bin/microgen
DATADIR=/tmp/microgen
BASEDIR=/home/mgrol/go/src/github.com/MarcGrol/microgen

nohup $MICROGEN -service=tour -port=8081 -base-dir=${DATADIR}/tour > $DATADIR/tour/log.txt &
nohup $MICROGEN -service=gambler -port=8082 -base-dir=${DATADIR}/gambler > $DATADIR/gambler/log.txt &
nohup $MICROGEN -service=news -port=8083 -base-dir=${DATADIR}/news > $DATADIR/news/log.txt &
nohup $MICROGEN -service=collector -port=8084 -base-dir=${DATADIR}/collector > $DATADIR/collector/log.txt &

nohup $MICROGEN -service=proxy -port=8080 -base-dir=${BASEDIR} &

ps -eaf|grep microgen |grep -v grep

