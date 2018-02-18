#!/bin/bash





for p in $(cat top1000english.txt); do
  printf 'http://localhost:8080/v1/global/%s\n' $p
done 

