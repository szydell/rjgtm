#!/bin/bash

export gtm_icu_version="5.0"
export gtm_chset="UTF-8"
export gtm_repl_instance="/home/szydell/.fis-gtm/V6.3-001A_x86_64/g/gtm.repl"
export gtm_chset="UTF-8"
export gtm_log="/tmp/fis-gtm/V6.3-001A_x86_64"
export gtm_prompt="GTM>"
export gtm_retention="42"
export gtmgbldir="/home/szydell/.fis-gtm/V6.3-001A_x86_64/g/gtm.gld"
export gtmroutines="/home/szydell/.fis-gtm/V6.3-001A_x86_64/o/utf8*(/home/szydell/.fis-gtm/V6.3-001A_x86_64/r /home/szydell/.fis-gtm/r) /opt/fis/6.3-001A/utf8/plugin/o/utf8(/opt/fis/6.3-001A/utf8/plugin/r) /opt/fis/6.3-001A/utf8/libgtmutil.so /opt/fis/6.3-001A/utf8 /home/szydell/code/src/gitlab.com/szydell/rjgtm/"
export gtmdir="/home/szydell/.fis-gtm"
export gtm_etrap='Write:(0=$STACK) "Error occurred: ",$ZStatus,!'
export gtm_principal_editing="EDITING"
export gtm_tmp="/tmp/fis-gtm/V6.3-001A_x86_64"
export gtm_dist="/opt/fis/6.3-001A/utf8"
export GTMCI="/home/szydell/code/src/gitlab.com/szydell/rjgtm/gtmacces.ci"
./rjgtm
