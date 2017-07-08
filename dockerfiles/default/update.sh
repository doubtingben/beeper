#!/bin/bash


SRCFILE=../../docker-compose.yml

function printlink() {
  echo $1
}

grep Host ${SRCFILE} | \
  sed "s/.*\"traefik.frontend.rule=Host:\([^']\+\)\".*/\1/" | \
  awk '{ printf( "%s \n", $1 ) }' | eval xargs printlink {} \;

