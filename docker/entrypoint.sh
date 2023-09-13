#!/sbin/tini /bin/sh

. /init &

sleep 20

cd /c8y-scanner && ./c8y-scanner
