#!/bin/sh --

export GOOS=linux
export CGO_ENABLED=0

find cmd -mindepth 1 -maxdepth 1 -type d | while read -r dir; do
  name=$(basename "${dir}")
  go build -o "build/$name/bootstrap" "./${dir}"
  mkdir -p "build/$name"
  cd "build/$name" || exit
  zip "../$name.zip" -r .
  cd ../.. || exit
done
