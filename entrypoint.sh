#!/bin/bash

if [[ $1 == "projects" ]]; then
  echo "Project create"
fi

atlas "$1" "$2"