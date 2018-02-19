#!/bin/bash


echo "Siege 1 url"
siege http://localhost:8080/v1/gvstats -r10 -c250 | grep -v "HTTP/1.1 200"

echo "Siege write 100 urls * 200 users"
siege -f siege_write_url.txt -ronce -c250 | grep -v "HTTP/1.1 200"

echo "Siege read 100 url * 200 users"
siege -f siege_read_url.txt -ronce -c250 | grep -v "HTTP/1.1 200"

echo "Siege data 100 urls * 200 users"
siege -f siege_data_url.txt -ronce -c250 | grep -v "HTTP/1.1 200"


#echo "Siege delete 100 urls * 200 users"
#siege -f siege_delete_url.txt -ronce -c250 | grep -v "HTTP/1.1 200"
