#!/bin/bash

export gtm_icu_version="5.0"
export gtm_chset="UTF-8"
export gtm_repl_instance="/home/szydell/.fis-gtm/V6.3-001A_x86_64/g/gtm.repl"
export gtm_chset="UTF-8"
export gtm_log="/tmp/fis-gtm/V6.3-001A_x86_64"
export gtm_prompt="GTM>"
export gtm_retention="42"
export gtmgbldir="/home/szydell/.fis-gtm/V6.3-001A_x86_64/g/gtm.gld"
export gtmroutines="/home/szydell/.fis-gtm/V6.3-001A_x86_64/o/utf8*(/home/szydell/.fis-gtm/V6.3-001A_x86_64/r /home/szydell/.fis-gtm/r) /opt/fis/6.3-001A/utf8/plugin/o/utf8(/opt/fis/6.3-001A/utf8/plugin/r) /opt/fis/6.3-001A/utf8/libgtmutil.so /opt/fis/6.3-001A/utf8 /home/szydell/code/src/github.com/szydell/rjgtm/m/"
export gtmdir="/home/szydell/.fis-gtm"
export gtm_tmp="/tmp/fis-gtm/V6.3-001A_x86_64"
export gtm_dist="/opt/fis/6.3-001A/utf8"
export GTMCI="/home/szydell/code/src/github.com/szydell/rjgtm/m/gtmacces.ci"
# enable mumps ^C
#export gtm_nocenable=0

export gtm_etrap='I 0=$ST W "Process terminated by: ",$ZS,! ZHALT 1'
export HUGETLB_SHM=no
./rjworker
