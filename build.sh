#! /bin/sh

set -e
registry=index.docker.io/djimydev


cd sa-frontend
docker build -t $registry/safrontend .
docker push $registry/safrontend
cd ..

cd sa-webapp
docker build -t $registry/sawebapp .
docker push $registry/sawebapp
cd ..

cd sa-logic
docker build -t $registry/salogic .
docker push $registry/salogic
cd ..