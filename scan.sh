#!/bin/bash
#base=${0##*/}

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
