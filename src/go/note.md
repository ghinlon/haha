# Note

# Links

# How to read code

* Overview First.

# Notes

1. package level only have: `const` `var` `func`;
   `expression` MUST put in `func`, so `expression` can put in `func init`.

1. always test a slice's `len`, don't test it's `s[x]`'s value, may out of
   index.

1. slice and map should `make` before use. or it'll `panic: assignment to entry
   in nil map`.
