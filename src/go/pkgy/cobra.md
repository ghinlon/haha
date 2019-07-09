# [cobra](https://godoc.org/github.com/spf13/cobra)

`import "github.com/spf13/cobra"`

# Links

* [GitHub - spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
* [Cobra Generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)
* [Building CLI tools in Go with Cobra - My Code Smells!](https://mycodesmells.com/post/building-cli-tools-in-go-with-cobra)
* [How to build great CLI’s in Golang, one snake at a time](https://medium.com/@skdomino/writing-better-clis-one-snake-at-a-time-d22e50e60056)

# Overview

1. `var rootCmd = &cobra.Command{}`
1. in `init()` do `cobra.OnInitialize(initConfig,y ...func())`, `rootCmd.AddCommand(cmds)`, `rootCmd.Flags().XXXVarP`, `viper.BindPFlag`, `viper.SetDefault`(may)
1. `func initConfig() { ...}`
1. `func Execute() { ... }`

# README

[cobra/README.md](https://github.com/spf13/cobra/blob/master/README.md)

# Getting Started

参考`cobra init`生成的[代码](https://github.com/iofxl/practisego/tree/master/08cobra/init)

in the project dir, run:

```
$ cobra init .
$ cobra add serve
$ cobra add config
$ cobra add create -p 'configCmd'
$ tree
.
├── LICENSE
├── cmd
│   ├── config.go
│   ├── create.go
│   ├── root.go
│   └── serve.go
└── main.go

// Now can run
$ go run main.go
$ go run main.go serve
$ go run main.go config
$ go run main.go config create
$ go run main.go help serve
```

# func OnInitialize

```go
func OnInitialize(y ...func())
```

OnInitialize sets the passed functions to be run when each command's Execute method is called.

# Command结构

rootCmd.AddCommand(cmds)要放在init()里面，程序的执行顺序`import --> const --> var --> init()`

**无法在包级别调用东西的方法**,所以*Command的其它的一些方法基本上统都要放在init()里去执行, 或者在自己的Run属性（类型`func(cmd *Command, args []string)`)里

```go
type Command struct {
    // Use is the one-line usage message.
    Use string
    ...
    // Run: Typically the actual work function. Most commands will only implement this.
    Run func(cmd *Command, args []string)
    ...
}

// Execute uses the args (os.Args[1:] by default)
// and run through the command tree finding appropriate matches
// for commands and then corresponding flags.
func (c *Command) Execute() error
// AddCommand adds one or more commands to this parent command.
func (c *Command) AddCommand(cmds ...*Command) {
    for i, x := range cmds {
        ...
	c.commands = append(c.commands, x)
	...
    }
}
// Commands returns a sorted slice of child commands.
func (c *Command) Commands() []*Command

func (c *Command) Flag(name string) (flag *flag.Flag)
func (c *Command) Flags() *flag.FlagSet
func (c *Command) PersistentFlags() *flag.FlagSet
```

# pflag.Flag结构和pflag.Value接口

Command结构通过Flag(name)方法返回Flag,然后Flag结构通过Value接口取得flag的值的字符串化

```go
type Flag struct {
	Name                string              // name as it appears on command line
	Shorthand           string              // one-letter abbreviated flag
	Usage               string              // help message
	Value               Value               // value as set
	DefValue            string              // default value (as text); for usage message
	Changed             bool                // If the user set the value (or if left to default)
	NoOptDefVal         string              // default value (as text); if the flag is on the command line without any options
	Deprecated          string              // If this flag is deprecated, this string is the new or now thing to use
	Hidden              bool                // used by cobra.Command to allow flags to be hidden from help/usage text
	ShorthandDeprecated string              // If the shorthand of this flag is deprecated, this string is the new or now thing to use
	Annotations         map[string][]string // used by cobra.Command bash autocomple code
}

type Value interface {
    String() string
    Set(string) error
    Type() string
    }

```

# pflag.FlagSet结构

A FlagSet represents a set of defined flags. 

`*Command.PersistentFlags()`方法和`*Command.Flags()`方法会返回这个结构

这个结构的各种方法，就是跟标准库里的flag.XXX函数作用一样

这个结构的`Lookup(name string) *Flag`方法返回*Flag结构

```go
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage func()

	// SortFlags is used to indicate, if user wants to have sorted flags in
	// help/usage messages.
	SortFlags bool

	// ParseErrorsWhitelist is used to configure a whitelist of errors
	ParseErrorsWhitelist ParseErrorsWhitelist
	// contains filtered or unexported fields
}

func (f *FlagSet) String(name string, value string, usage string) *string
func (f *FlagSet) StringP(name, shorthand string, value string, usage string) *string
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
func (f *FlagSet) StringVarP(p *string, name, shorthand string, value string, usage string)

func (f *FlagSet) Lookup(name string) *Flag

// viper.BindPFlag
// BindPFlag binds a specific key to a pflag (as used by cobra).
// Example (where serverCmd is a Cobra instance):
//
//   serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
//   Viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
//
func BindPFlag(key string, flag *pflag.Flag) error { return v.BindPFlag(key, flag) }
```
Example:

[create-rootcmd](https://github.com/spf13/cobra/#create-rootcmd)

```go
func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
  rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
  rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
  rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
  rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

  viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
  viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
  viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

  viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
  viper.SetDefault("license", "apache")
}
```

# Working with Flags

## Assign flags to a command

Since the flags are defined and used in different locations, we need to define a variable outside with the correct scope to assign the flag to work with.

```go
var Verbose bool
var Source string
```

There are two different approaches to assign a flag.

### Persistent Flags

A flag can be 'persistent' meaning that this flag will be available to the command it's assigned to as well as every command under that command. For global flags, assign a flag as a persistent flag on the root.

`rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")`

### Local Flags

A flag can also be assigned locally which will only apply to that specific command.

`rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")`

## Local Flag on Parent Commands

By default Cobra only parses local flags on the target command, any local flags on parent commands are ignored. By enabling Command.TraverseChildren Cobra will parse local flags on each command before executing the target command.

```go
command := cobra.Command{
  Use: "print [OPTIONS] [COMMANDS]",
    TraverseChildren: true,
}
```

## Bind Flags with Config

You can also bind your flags with viper:

```go
var author string

func init() {
  rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
    viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}
```
In this example the persistent flag `author` is bound with `viper`. **Note**, that the variable `author` will not be set to the value from config, when the `--author` flag is not provided by user.

## Required flags

Flags are optional by default. If instead you wish your command to report an error when a flag has not been set, mark it as required:

```go
rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
rootCmd.MarkFlagRequired("region")
```
