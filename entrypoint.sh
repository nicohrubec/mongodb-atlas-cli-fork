#!/bin/bash

if [[ $1 == "projects" && $2 == "create" ]]; then
  projectId=(atlas projects create "$3" -o json-path="$.id")
  echo "projectId=$projectId" >> $GITHUB_OUTPUT
else
  atlas "$1" "$2" "$3"
fi