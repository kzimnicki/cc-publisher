#!/bin/bash


echo "$(date) start" >> /tmp/startup.log
cd /home/pi/rpi/zvbi-0.2.35/test;

echo "run tzap" -c /home/pi/.tzap/channels.conf >> /tmp/startup.log
tzap "ZDF" -c /home/pi/.tzap/channels.conf 2> /dev/null &

echo "sleep" >> /tmp/startup.log
sleep 10

echo "capture cc" >> /tmp/startup.log

./capture --device /dev/dvb/adapter0/demux0 --pid 551 --sliced | ./ttxfilter 777 | tee /media/PENDRIVE/rpi/zdf.cc.bin | ../../../go/bin/go run CCPublisher.go 2> /media/PENDRIVE/rpi/error.log 
echo "end" >> /tmp/startup.log

