# awk

* use system command  
  `awk '/^VoiceRoot/ { if ( ! system ( "test -d " $2 ) )  system("ln -sf " $2" " $1 ) }' <file>`
* use shell variable as pattern  
  `a="patter1"; awk -v a="$a" '$0 ~ a { print }'`

