#!/bin/bash





for p in $(cat top1000english.txt); do
  printf 'http://localhost:8080/v1/data/%s\n' $p
done 

