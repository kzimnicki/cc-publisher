#!/bin/bash

log() {  # classic logger 
   local prefix="[$(date +%Y/%m/%d\ %H:%M:%S)]: "
   echo "${prefix} $@" >&2
} 

GO_HOME="/home/pi/go/bin"

log "start"

cd ${0%/*}

log `pwd`

log "run tzap"
./rpi_dvbt_tools/tzap "ZDF" -c rpi_dvbt_tools/channels.conf 2> /dev/null &

log "sleep 10 seconds"
sleep 10

log "capture cc"

./rpi_dvbt_tools/capture --device /dev/dvb/adapter0/demux0 --pid 551 --sliced | ./rpi_dvbt_tools/ttxfilter 777 | tee -a /media/PENDRIVE/rpi/zdf.cc.bin | $GO_HOME/go run CCPublisher.go >> /media/PENDRIVE/rpi/error.log 

log "end"

sleep 5
