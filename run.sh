#!/bin/sh

microgen -service=proxy -port=8080 &
microgen -service=tour -port=8081 &
microgen -service=gambler -port=8082 &
microgen -service=collector -port=8084 &
