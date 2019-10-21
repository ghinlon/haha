# Principle

# Links

# How to read code

* Don't follow functions always. make a bird view first.

# principle

1. package level only have: `const var func`, `expression` MUST put in `func`,
   so can put in `func init`.

1. always test a slice with `len(s) > N` first, don't test slice with `s[n-1] == XXX`, it may out of index.

1. be carefull of `panic: assignment to entry in nil map`, assign value to
   slice and map need to `make` it first.

# What is function

1. make up some values, manipulate the values, and return the values.

1. a function is combinations of functiongs.



