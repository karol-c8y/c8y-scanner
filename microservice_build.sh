#!/usr/bin/bash

rm microservice/c8y-scanner.zip

./build.sh

docker save c8y-scanner -o microservice/image.tar

cd microservice
zip c8y-scanner.zip image.tar cumulocity.json
cd ..

rm microservice/image.tar
