#!/bin/bash

docker kill $(docker ps -q -f ancestor=c8y-scanner)

docker run -it \
	-p 8888:80 \
	-e "C8Y_BASEURL=https://c8y-scanner.latest.stage.c8y.io" \
	-e "C8Y_BOOTSTRAP_TENANT=t177577519" \
	-e "C8Y_BOOTSTRAP_USER=servicebootstrap_c8y-scanner" \
	-e "C8Y_BOOTSTRAP_PASSWORD=0WNW2loQ0Y2CVfOjJ9d3Uwdun2o99JUa" \
	c8y-scanner


