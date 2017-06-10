#!/bin/bash

if [ $# -eq 0 ]
then
  echo "Sites parameter is required"
  exit 1
fi

curl -XPOST -s http://localhost:8080/check -d @$1
