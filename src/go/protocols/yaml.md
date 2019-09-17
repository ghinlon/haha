# YAML

YAML stands for "YAML Ain't Markup Language" 

# Links

* [YAML Syntax | Grav Documentation](https://learn.getgrav.org/advanced/yaml)
* [YAML and Maps in Go | □▢ SQUARISM](https://squarism.com/2014/10/13/yaml-go/)

# Basic Rules


* YAML files should end in .yaml whenever possible in Grav.
* YAML is case sensitive.
* YAML **does not allow the use of tabs**. Spaces are used instead as tabs are not universally supported.

# Basic Data Types

YAML excels at working with **mappings** (hashes / dictionaries), **sequences** (arrays / lists), and **scalars** (strings / numbers). 

# Scalars

```
integer: 25
string: "25"
float: 25.0
boolean: Yes
```

# Sequences

```
- Cat
- Dog
- Goldfish
```

**nested sequence:**

```
-
 - Cat
 - Dog
 - Goldfish
-
 - Python
 - Lion
 - Tiger
```

# Mappings

* Block list items include same indentation as the surrounding block level
  because `-` is considered as a part of indentation.

```
animal: pets
```

```
pets:
 - Cat
 - Dog
 - Goldfish
```

## Tricky

```
fruits:
  apple:
    - red

------------------
Fruits map[string][]string
// Fruits have properties.  You have success.
```

```
fruits:
  - apple:
    - red

------------------
Fruits []map[string][]string
// There are more fruits now.  Good job, fruit adventurer.
```

This is because apple has become an array value and not a key. 

In other words, we can have many apples in the second example but not the first. 

为什么第一个不是：

```go
type Anything struct {
	Fruits map[string][]string
}
```

而第二个不是：

```go
type Anything struct {
	Fruits []map[string][]string
}
```

## An Example

```go
type A struct {
        B string
        C int
}

type AS struct {
        As
}

type As []A

func main() {

        var as = As{{"zhang1", 1}, {"zhang2", 2}, {"zhang3", 3}}
        var aS = AS{As{{"zhang1", 1}, {"zhang2", 2}, {"zhang3", 3}}}

        b, err := yaml.Marshal(as)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(string(b))
        fmt.Println()

        bS, err := yaml.Marshal(aS)

        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(string(bS))

}
/* output:
- b: zhang1
  c: 1
- b: zhang2
  c: 2
- b: zhang3
  c: 3


as:
- b: zhang1
  c: 1
- b: zhang2
  c: 2
- b: zhang3
  c: 3
*/
```

