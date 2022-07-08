#!/usr/bin/env bash

c-for-go goplctag.yml
c-for-go callbacks.yml
cp callbacks.h callbacks/
cp callbacks.c callbacks/