#!/usr/bin/env bash

c-for-go -nostamp goplctag.yml
cp -R ./goplctag/* ./
rm -r ./goplctag/