#!/bin/sh

nohup nsqlookupd > /tmp/nsqlookupd.log & 
nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 > /tmp/nsq.log &
nohup nsqadmin --lookupd-http-address=127.0.0.1:4161 > /tmp/nsqadmin.log &


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
