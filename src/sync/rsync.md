# rsync

# Links

* [How To Use Rsync to Sync Local and Remote Directories on a VPS | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-use-rsync-to-sync-local-and-remote-directories-on-a-vps)

# Basic

sync the content of dir1 to dir2

```
rsync -av dir1/ dir2
rsync -avhW --progress /src/ /dst/
```

```
-a is for archive, which preserves ownership, permissions etc.
-v is for verbose, so I can see what's happening (optional)
-h is for human-readable, so the transfer rate and file sizes are easier to read (optional)
-W is for copying whole files only, without delta-xfer algorithm which should reduce CPU load
--progress so I can see the progress of large files (optional)
```

By default, rsync does not delete anything from the destination directory.

We can change this behavior with the `--delete` option. Before using this
option, use the `--dry-run` option and do testing to prevent data loss:

```
rsync -a --delete source destination
```

# Exclude Files

[6 rsync Examples to Exclude Multiple Files and Directories using exclude-from](https://www.thegeekstuff.com/2011/01/rsync-exclude-files-and-folders/?utm_source=feedburner)

* Exclude path is always relative

```
rsync -avz --exclude 'dir1' --exclude '*.foo' source/ destination/
// or
rsync -avz --exclude-from 'exclude-list.txt' source/ destination/

cat exclude-list.txt
dir1
*.foo

     --exclude=PATTERN       exclude files matching PATTERN
     --exclude-from=FILE     read exclude patterns from FILE
     --include=PATTERN       don't exclude files matching PATTERN
     --include-from=FILE     read include patterns from FILE
     --files-from=FILE       read list of source-file names from FILE
```



