# sed
  
* Sed修改某一行  
  `sed '5s/^.*$/xxxxx/'  file`
* sed修改下一行  
  `sed '/DIVIDER/{n;s/.*/[begin]&[end]\n/;}' file1`
* sed打印下一行  
  `sed -n '/^\[dns/{n;p}' /etc/cobbler/modules.conf`
* sed打印本行加下一行  
  `sed -n '/^\[dns/{N;p}' /etc/cobbler/modules.conf`


