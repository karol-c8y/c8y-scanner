#!/bin/bash

docker kill $(docker ps -q -f ancestor=c8y-scanner)

docker run -p 8888:8888 -d c8y-scanner
