#! /bin/sh

set -e

cd sa-frontend  && 
docker build -t localhost:8888/safrontend .  && 
docker push localhost:8888/safrontend  && 
cd ..  &&  

cd sa-webapp && 
docker build -t localhost:8888/sawebapp . && 
docker push localhost:8888/sawebapp  && 
cd ..  && 

cd sa-logic  && 
docker build -t localhost:8888/salogic .  && 
docker push localhost:8888/salogic  && 
cd ..