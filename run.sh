#!/bin/sh

pkill microgen

find . -name "*.db" -exec rm -f {} \;

microgen -service=proxy -port=8080 &
microgen -service=tour -port=8081 &
microgen -service=gambler -port=8082 &
microgen -service=news -port=8083 &
microgen -service=collector -port=8084 &

ps -eaf|grep microgen |grep -v grep


# boot2docker stop
# boot2docker download
# boot2docker up
# docker build -t microgen .
# docker run -ti -p 8081:8081 microgen /go/bin/microgen -service=tour -port=8081

# docker run --name lookupd -p 4160:4160 -p 4161:4161 nsqio/nsq /nsqlookupd

# docker run --name nsqd -p 4150:4150 -p 4151:4151 \
#    nsqio/nsq /nsqd \
#   --broadcast-address=172.17.42.1 \
#    --lookupd-tcp-address=172.17.42.1:4160


# docker-compose up