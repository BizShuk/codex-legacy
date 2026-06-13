#!/bin/bash


cat words.txt | tr " " "\n" | grep -v "^\s*$" | sort | uniq -c | sort -r | awk '{print $2 " " $1}'


