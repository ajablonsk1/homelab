#!/bin/bash

FOLDER=$1
FILENAME=$2

cd "$FOLDER" && nvim "$FILENAME"
