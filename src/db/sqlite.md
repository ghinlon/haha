# [SQLite](https://sqlite.org/index.html)

# Links

* [Command Line Shell For SQLite](https://sqlite.org/cli.html)

# Example

```
$ sqlite3 ex1
SQLite version 3.28.0 2019-03-02 15:25:24
Enter ".help" for usage hints.
sqlite> create table tbl1(one varchar(10), two smallint);
sqlite> insert into tbl1 values('hello!',10);
sqlite> insert into tbl1 values('goodbye', 20);
sqlite> select * from tbl1;
hello!|10
goodbye|20
sqlite>
```

# Basics

```
.tables ?TABLE?          List names of tables matching LIKE pattern TABLE
.schema ?PATTERN?        Show the CREATE statements matching PATTERN

.schema Album 		
pragma table_info(Album); 		// get column information of a table.
```

# Backup

## archive and reconstruct

The text format is pure SQL so you can also use the `.dump` command to export
an SQLite database into other popular SQL database engines.

```
// dump
sqlite3 ex1 .dump | gzip -c >ex1.dump.gz
// reconstruct
zcat ex1.dump.gz | sqlite3 ex2
```

# Misc

## Get the last record


```
SELECT * FROM table ORDER BY column DESC LIMIT 1;
```

# sqlite3 driver

* [sqlite3 - GoDoc](https://godoc.org/github.com/mattn/go-sqlite3#SQLiteDriver.Open)



