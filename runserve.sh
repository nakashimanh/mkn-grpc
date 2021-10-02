#!/bin/bash

# scan docker images
DIR="docker"
if [ -d $DIR ];then
  for entry in "$DIR"/*
  do
    first_line=$(head -n 1 "$entry")
    image="${first_line/"FROM "/}"
    docker scan --file "$entry" "$image"
  done
else
  echo "Directory not exists."
  exit 1
fi

docker-compose up --build
#docker-compose exec grpc sh
docker-compose down