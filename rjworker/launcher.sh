#!/bin/bash
export gtm_dist="/opt/fis-gtm/6.3-001A/utf8"
export gtm_chset="UTF-8"
. /opt/fis-gtm/6.3-001A/gtmprofile
export gtm_icu_version="5.0"
export gtm_repl_instance="/home/${USER}/.fis-gtm/V6.3-001A_x86_64/g/gtm.repl"

export gtm_log="/tmp/fis-gtm/V6.3-001A_x86_64"
export gtm_prompt="GTM>"
export gtm_retention="42"
export gtmgbldir="/home/${USER}/.fis-gtm/V6.3-001A_x86_64/g/gtm.gld"
#export gtmroutines="/home/${USER}/.fis-gtm/V6.3-001A_x86_64/o/utf8*(/home/${USER}/.fis-gtm/V6.3-001A_x86_64/r /home/${USER}/.fis-gtm/r) /opt/fis/6.3-001A/utf8/plugin/o/utf8(/opt/fis/6.3-001A/utf8/plugin/r) /opt/fis/6.3-001A/utf8/libgtmutil.so /opt/fis/6.3-001A/utf8 /home/${USER}/code/src/github.com/${USER}/rjgtm/m/"
export gtmdir="/home/${USER}/.fis-gtm"
export gtm_tmp="/tmp/fis-gtm/V6.3-001A_x86_64"
#export GTMCI="/home/${USER}/code/src/github.com/${USER}/rjgtm/m/gtmacces.ci"
#enable mumps ^C
#export gtm_nocenable=0

export gtm_etrap='I 0=$ST W "Process terminated by: ",$ZS,! ZHALT 1'
export HUGETLB_SHM=no
#./rjworker
