#!/bin/bash
#########################################################################################
#											
#											
#											
########################################################################################## 

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)

EVENTS_NUMBERS=$1

for (( i=1; i <= ${EVENTS_NUMBERS}; i++ ))
do
 sleep 1
 echo "Random number $i"
 DATA="Data number $i"
 echo "db.events.save({'transport':'f1com','type':1, 'code': 'NODSTA', 'data':'${DATA}', 'appid': 'builkgenerator','asteriskid':'astersk1'})" | mongo 192.168.3.20/notifications
 
done

#echo "db.events.findOne()" | mongo notifications
#echo "show collections" | mongo notifications
#mongoimport --db notifications --collection events --file ${SCRIPTPATH}/events.json