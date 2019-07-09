# simple anonymous ftp server

# Links

* [Set up an anonymous FTP server with vsftpd in less than a minute | G-Loaded Journal](https://www.g-loaded.eu/2008/12/02/set-up-an-anonymous-ftp-server-with-vsftpd-in-less-than-a-minute/)

Config 

rhel6: `/etc/vsftpd/vsftpd.conf`

```
use_localtime=YES

local_enable=NO
anonymous_enable=YES
no_anon_password=YES
write_enable=NO

anon_root=/ftp_root_dir/
```

Run

`/etc/init.d/vsftpd start`

# Multi Dir

comment  `anon_root`, so to se the default ftp dir `/var/ftp/pub/`

then config `mount --bind` and put them in `/etc/rc.local`

```
mount --bind /dir1/ /var/ftp/pub/dir1
mount --bind /dir2/ /var/ftp/pub/dir2
```

# log

`xferlog_enable` control whether log or not.

```
# The target log file can be vsftpd_log_file or xferlog_file.
# This depends on setting xferlog_std_format parameter
xferlog_enable=YES
```

`xferlog_std_format` is a switch to control which file to log

```
# Switches between logging into vsftpd_log_file and xferlog_file files.
# NO writes to vsftpd_log_file, YES to xferlog_file
xferlog_std_format=YES
```

if `xferlog_std_format=NO`, then need to config this:

```
vsftpd_log_file=/var/log/vsftpd.log
log_ftp_protocol=YES
```



