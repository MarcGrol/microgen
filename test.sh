#!/bin/sh

HOSTNAME="192.168.59.103"
curl -X POST --header "Content-type: application/json"  --header "Accept: application/json" --data '{"year":2015}' "http://${HOSTNAME}:8081/api/tour"
curl -X GET  --header "Accept: application/json" "http://${HOSTNAME}:8081/api/tour/2015"
curl -X POST --header "Content-type: application/json"  --header "Accept: application/json" --data '{"year":2015,"id":9,"name":"VAN GARDEREN Tejay","team":"BMC"}}' "http://${HOSTNAME}:8081/api/tour/2015/cyclist"


