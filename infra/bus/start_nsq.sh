#!/bin/sh -x

MICROGEN_ROOT=${GOPATH}/src/github.com/MarcGrol/microgen/

mkdir -p ${MICROGEN_ROOT}/log
nohup nsqlookupd > ${MICROGEN_ROOT}/log/nsqlookupd.log & 
nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 > ${MICROGEN_ROOT}/log/nsq.log &
nohup nsqadmin --lookupd-http-address=127.0.0.1:4161 > ${MICROGEN_ROOT}/log/nsqadmin.log &

sleep 1
ps -eaf|grep nsq

# create topics if needed

curl -X POST localhost:4161/topic/create?topic=tourApp_TourCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_EtappeCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_CyclistCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_EtappeResultsCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_GamblerCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_GamblerTeamCreated
curl -X POST localhost:4161/topic/create?topic=tourApp_CyclistScoreCalculated
curl -X POST localhost:4161/topic/create?topic=tourApp_GamblerScoreCalculated
curl -X POST localhost:4161/topic/create?topic=tourApp_NewsItemCreated

#
# stats
#
# open url http://localhost:4171/ in browser
# nsq_stat --topic=mytopic --channel=ch -lookupd-http-address 127.0.0.1:4161

#
# producers
#
# producer
# curl -d 'hello world 1' 'http://127.0.0.1:4151/put?topic=mytop'c'

#
# consumer
#

# consumer
# nsq_tail --topic=mytopic --lookupd-http-address=127.0.0.1:4161
# nsq_to_file --topic=test --output-dir=. --lookupd-http-address=127.0.0.1:4161


# webinterface see http://localhost:4171/
