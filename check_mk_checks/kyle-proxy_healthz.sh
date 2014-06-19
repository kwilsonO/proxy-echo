#!/bin/bash
OK=0
WARN=1
CRIT=2
UNKNOWN=3
# the service name needs the container id appended to it
SERVICE=kyle-proxy_healthz_$2
EXPECTED="OK"

# The port needs to be read from the command line because it is assigned dynamically by the supervisor
OUTPUT=`curl -s http://127.0.0.1:$1/healthz`

if [ "$OUTPUT" = "OK" ] ; then
  echo $OK $SERVICE - /healthz seems OK OUTPUT: $OUTPUT
else
  echo $CRIT $SERVICE - /healthz returned error OUTPUT: $OUTPUT
fi
