#!/bin/bash

if [[ $1 == "projects" && $2 == "create" ]]; then
  atlas projects create "$3"
else
  atlas "$1" "$2" "$3"
fi