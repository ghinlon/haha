# Bash

# Links

# what is $$

`$$` is the process ID (PID) of the script itself. 

# Execute binary from curl

[bash - execute binary directly from curl - Stack Overflow](https://stackoverflow.com/questions/14909827/execute-binary-directly-from-curl/14909927#14909927)

```sh
#!/bin/sh

# Arrange for the temporary file to be deleted when the script terminates
trap 'rm -f "/tmp/exec.$$"' 0
trap 'exit $?' 1 2 3 15

# Create temporary file from the standard input
cat >/tmp/exec.$$

# Make the temporary file executable
chmod +x /tmp/exec.$$

# Execute the temporary file
/tmp/exec.$$
```

# Read Line

[bash - Read a file line by line assigning the value to a variable - Stack Overflow](https://stackoverflow.com/questions/10929453/read-a-file-line-by-line-assigning-the-value-to-a-variable/10929511#10929511)

```bash
#!/bin/bash
while IFS='' read -r line || [[ -n "$line" ]]; do
    echo "Text read from file: $line"
done < "$1"
```

Explanation:

* `IFS=''` (or `IFS=`) prevents leading/trailing whitespace from being trimmed.
* `-r` prevents backslash escapes from being interpreted.
* `|| [[ -n $line ]]` prevents the last line from being ignored if it doesn't
  end with a `\n` (since `read` returns a non-zero exit code when it encounters
  EOF).

# Math calculation

[Performing Math calculation in Bash - Shell Tips!](https://www.shell-tips.com/2010/06/14/performing-math-calculation-in-bash/)

* `myvar=$(expr 1 + 1)`
* `let myvar+=1`
* `myvar=$((myvar+3))`


