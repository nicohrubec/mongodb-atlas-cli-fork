#!/bin/bash

if [[ $1 == "projects" && $2 == "create" ]]; then
  echo "$1"
  echo "$2"
  echo "$3"
  atlas projects create "$3"
else
  atlas "$1" "$2" "$3"
fi