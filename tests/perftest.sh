#!/bin/bash
echo "----------------------------------------------------"
echo " strOrSlice test :"
time ./strOrSlice 1>/dev/null
./xtime.sh ./strOrSlice 1>/dev/null
echo "----------------------------------------------------"
echo " strOrSlice2 test :"
time ./strOrSlice2 1>/dev/null
./xtime.sh ./strOrSlice2 1>/dev/null

