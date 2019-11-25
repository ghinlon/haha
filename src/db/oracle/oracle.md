# oracle

# Links

* 

# Overview

# Create listener

* [ORACLE-BASE - Oracle Network Configuration (listener.ora , tnsnames.ora , sqlnet.ora)](https://oracle-base.com/articles/misc/oracle-network-configuration)

```
cat > listener.ora <<EOF
LISTENER =
  (DESCRIPTION_LIST =
    (DESCRIPTION =
      (ADDRESS = (PROTOCOL = IPC)(KEY = EXTPROC1)) 
      (ADDRESS = (PROTOCOL = TCP)(HOST = myserver.example.com)(PORT = 1521))
    )
  )

SID_LIST_LISTENER =
  (SID_LIST =
    (SID_DESC =
      (GLOBAL_DBNAME = orcl.example.com)
      (ORACLE_HOME = /u01/app/oracle/product/11.2.0.4/db_1)
      (SID_NAME = orcl)
    )
  )
EOF
```


