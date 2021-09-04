#!/bin/bash
python3 main.py --text "$@" | tr '\n' ' ' | sed 's/.*=//' | tr " " "\n" | sort | uniq
