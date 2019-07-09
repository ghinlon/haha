# Wiki

# Links

* [CodeReviewComments Â· golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments)

1. package level only have: `const var func`, `expression` MUST put in `func`,
   so can put in `func init`.
1. always test a slice with `len(s) > N` first, don't test slice with `s[n-1]
   == XXX`, it may out of index.
1. be carefull of `panic: assignment to entry in nil map`, assign value to
   slice and map need to `make` it first.

# Hex

```
0+0=0                                                                                                                                                                                  
1+0=1   1+1=2                                                                                                                                                                          
2+0=2   2+1=3   2+2=4                                                                                                                                                                  
3+0=3   3+1=4   3+2=5   3+3=6                                                                                                                                                          
4+0=4   4+1=5   4+2=6   4+3=7   4+4=8                                                                                                                                                  
5+0=5   5+1=6   5+2=7   5+3=8   5+4=9   5+5=a                                                                                                                                          
6+0=6   6+1=7   6+2=8   6+3=9   6+4=a   6+5=b   6+6=c                                                                                                                                  
7+0=7   7+1=8   7+2=9   7+3=a   7+4=b   7+5=c   7+6=d   7+7=e                                                                                                                          
8+0=8   8+1=9   8+2=a   8+3=b   8+4=c   8+5=d   8+6=e   8+7=f   8+8=10                                                                                                                 
9+0=9   9+1=a   9+2=b   9+3=c   9+4=d   9+5=e   9+6=f   9+7=10  9+8=11  9+9=12                                                                                                         
a+0=a   a+1=b   a+2=c   a+3=d   a+4=e   a+5=f   a+6=10  a+7=11  a+8=12  a+9=13  a+a=14                                                                                                 
b+0=b   b+1=c   b+2=d   b+3=e   b+4=f   b+5=10  b+6=11  b+7=12  b+8=13  b+9=14  b+a=15  b+b=16                                                                                         
c+0=c   c+1=d   c+2=e   c+3=f   c+4=10  c+5=11  c+6=12  c+7=13  c+8=14  c+9=15  c+a=16  c+b=17  c+c=18                                                                                 
d+0=d   d+1=e   d+2=f   d+3=10  d+4=11  d+5=12  d+6=13  d+7=14  d+8=15  d+9=16  d+a=17  d+b=18  d+c=19  d+d=1a                                                                         
e+0=e   e+1=f   e+2=10  e+3=11  e+4=12  e+5=13  e+6=14  e+7=15  e+8=16  e+9=17  e+a=18  e+b=19  e+c=1a  e+d=1b  e+e=1c                                                                 
f+0=f   f+1=10  f+2=11  f+3=12  f+4=13  f+5=14  f+6=15  f+7=16  f+8=17  f+9=18  f+a=19  f+b=1a  f+c=1b  f+d=1c  f+e=1d  f+f=1e 
```

a = 10进制中的10，但a就是16进制中的a. 9+9=18的意义是进一位余8，a+a=14的意思是进
1位余4

# Args

**package os**

```go
var Args []string
    Args hold the command-line arguments, starting with the program name.
```

**package flag**

```go
// Args returns the non-flag command-line arguments.
func Args() []string

// Arg returns the i'th command-line argument. Arg(0) is the first remaining
// argument after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func Arg(i int) string

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int
```

**package pflag**

```go
func Args() []string
func Arg(i int) string
func NArg() int
```

**package cobra**

```go
// index args withs args in Run
type Command struct {
    // Run: Typically the actual work function. Most commands will only
    implement this.
    Run func(cmd *Command, args []string)

    // ... other fields elided ...
}
```

* [How to check if a file exists in Go? - Stack Overflow](how_to_check_if_a_file_exists.md)
* [distinguish ipv4 and ipv6](distinguish_ipv4_and_ipv6.md)
* [gorutine and CWD](gorutine_and_cwd.md)

