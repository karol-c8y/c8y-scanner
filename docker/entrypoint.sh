#!/sbin/tini /bin/sh

. /init &

cd /c8y-scanner && ./c8y-scanner
