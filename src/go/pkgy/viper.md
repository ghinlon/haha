# [viper](https://godoc.org/github.com/spf13/viper)

`import "github.com/spf13/viper"`

# Links

* [GitHub - spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)
* [Creating a Microservice Boilerplate in Go | Netlify](https://www.netlify.com/blog/2016/09/06/creating-a-microservice-boilerplate-in-go/)

`go get -u github.com/spf13/viper`

# Overview

# Tips

1. `func ReadInConfig` 去加载配置文件，在这之前或者没找到配置文件，`func ConfigFileUsed` 返回都是空的。
2. 只要`viper.BindPFlag`过，不管有无读到配置文件，`viper.GetXX`,`viper.Unmarshal(&cfg)`统都有法使的
3. 如果没读得配置文件，有法使`viper.WriteConfigAs("app.yaml")`保存
4. `viper.Unmarshal(&v)` can't directly unmarshal into a slice.

优先级：

flag设置 > 配置文件设置 > `viper.SetDefault` > flag默认值

# README

* [viper/README.md](https://github.com/spf13/viper/blob/master/README.md)

# SetConfgName and SetConfigType

SetConfgName不需要包含扩展名

```go
// SetConfigFile explicitly defines the path, name and extension of the config file.
// Viper will use this and not check any of the config paths.
func SetConfigFile(in string)
func SetConfigName(in string)
func SetConfigType(in string)
```

＃ AddConfigPath

可以执行多次，也就是可以添加多个目录

```go
func AddConfigPath(in string)
```

例：

```go
viper.SetConfigName("config") // name of config file (without extension)
viper.SetConfigType("yaml")

viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
viper.AddConfigPath("/usr/local/etc/appname/")   // path to look for the config file in
viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
viper.AddConfigPath(".")               // optionally look for config in the working directory

err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```

# ReadInConfig and Reading Config from io.Reader

```go
func ReadInConfig() error
func ReadConfig(in io.Reader) error
```

# Getting Values From Viper

Get returns an interface. For a specific value use one of the Get____ methods. 

```go
func Get(key string) interface{}
func GetBool(key string) bool

func InConfig(key string) bool
func IsSet(key string) bool
```

#  Unmarshaling a Config struct

```go
func Unmarshal(rawVal interface{}, opts ...DecoderConfigOption) error
```

Example:

```go
// in conf/config.go
type Config struct {
	Port string
}

func LoadConfig()(*Config, error) {
// all the other loading code

	config := new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

// all the error signatures above had to change to nil, err
	return config, nil
}
```


# WriteConfig

```go
func WriteConfig() error
func WriteConfigAs(filename string) error
// SafeWriteConfig writes current configuration to file only if the file does not exist.
func SafeWriteConfig() error
func SafeWriteConfigAs(filename string) error

// ConfigFileUsed returns the file used to populate the config registry.
func ConfigFileUsed() string
```

# Viper结构

```go
type Viper struct {
    // contains filtered or unexported fields
}

func GetViper() *Viper
// New returns an initialized Viper instance. 
func New() *Viper
func Sub(key string) *Viper
```

# Make viper connection with Variable

```go
func Set(key string, value interface{})
```

Example:

```go
type Config struct {
	Port int
}

var cfg Config

func initConfig() {
	viper.Set("secret", cfg.Port)
}
```

