#!/usr/bin/env sh

OLD_PWD=$(pwd)
cd $(dirname $0)

go run . $1 inputs/d$(printf %02d $1)/input.txt

cd $OLD_PWD