# grep
  
* 从文件尾开始匹配  
  ```
tac foo | grep -m 3 foo
-m 匹配到3次后就停止读文件,这样会非常效率.
```
* 高亮显示匹配  
  `--color=auto`


