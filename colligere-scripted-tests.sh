#!/usr/bin/env bash

rm -f colligere

go build

./colligere

./colligere config

./colligere config create

./colligere config view

./colligere config set

./colligere config view

# eventually create a fake URI localhost to test against. Not sure what to use, how to do it, etc, but will determine.

echo ./colligere deluge

rm -f colligere