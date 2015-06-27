#!/bin/sh

MICROGEN=/home/mgrol/go/bin/microgen
DATADIR=/tmp/microgen
BASEDIR=/home/mgrol/go/src/github.com/MarcGrol/microgen

nohup $MICROGEN -service=tour -port=8081 -base-dir=${DATADIR}/tour &
nohup $MICROGEN -service=gambler -port=8082 -base-dir=${DATADIR}/gambler &
nohup $MICROGEN -service=news -port=8083 -base-dir=${DATADIR}/news &
nohup $MICROGEN -service=collector -port=8084 -base-dir=${DATADIR}/collector &

nohup $MICROGEN -service=proxy -port=8080 -base-dir=${BASEDIR} &

ps -eaf|grep microgen |grep -v grep


# boot2docker stop
# boot2docker download
# boot2docker up
# docker build  --no-cache -t microgen .
# docker run -ti -p 8081:8081 microgen /go/bin/microgen -service=tour -port=8081

# docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

# docker run --name nsqd -p 4150:4150 -p 4151:4151 \
#    nsqio/nsq /nsqd \
#   --broadcast-address=172.17.42.1 \
#    --lookupd-tcp-address=172.17.42.1:4160


# docker-compose up

#nohup docker run -ti -p 8081:8080 microgen /go/bin/microgen -service=proxy -port=8080 -address=92.168.59.103 -base-dir=/go/src/github.com/MarcGrol/microgen&
#nohup docker run -ti -p 8081:8081 microgen /go/bin/microgen -service=tour -port=8081 -address=10.0.2.15 -base-dir=. &
#nohup docker run -ti -p 8081:8082 microgen /go/bin/microgen -service=gambler -port=8082 -address=92.168.59.103 -base-dir=. &
#nohup docker run -ti -p 8081:8083 microgen /go/bin/microgen -service=news -port=8083 -address=92.168.59.103 -base-dir=. &
#nohup docker run -ti -p 8081:8084 microgen /go/bin/microgen -service=collector -port=8084 -address=92.168.59.103 -base-dir=. &
