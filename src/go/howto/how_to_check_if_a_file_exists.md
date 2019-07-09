# How to check if a file exists ?

# Links

* [How to check if a file exists in Go? - Stack Overflow](https://stackoverflow.com/a/22483001)

```
Answer by Caleb Spare posted in gonuts mailing list.

    [...] It's not actually needed very often and [...] using os.Stat is easy enough for the cases where it is required.

    [...] For instance: if you are going to open the file, there's no reason to check whether it exists first. The file could disappear in between checking and opening, and anyway you'll need to check the os.Open error regardless. So you simply call os.IsNotExist(err) after you try to open the file, and deal with its non-existence there (if that requires special handling).

    [...] You don't need to check for the paths existing at all (and you shouldn't).

        os.MkdirAll works whether or not the paths already exist. (Also you need to check the error from that call.)

        Instead of using os.Create, you should use os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666) . That way you'll get an error if the file already exists. Also this doesn't have a race condition with something else making the file, unlike your version which checks for existence beforehand.

Taken from: https://groups.google.com/forum/#!msg/golang-nuts/Ayx-BMNdMFo/4rL8FFHr8v4J
```


