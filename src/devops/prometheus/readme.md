# [prometheus - GoDoc](https://godoc.org/github.com/prometheus/prometheus)

# Links

* [prometheus/internal_architecture.md at master Â· prometheus/prometheus](https://github.com/prometheus/prometheus/blob/master/documentation/internal_architecture.md)
* [github.com/oklog/run - GoDoc](https://godoc.org/github.com/oklog/run)

# Actors

```go
var g run.Group

g.Add( func() error {...}, func( err error ) {...} )
g.Add( func() error {...}, func( err error ) {...} )
...
err := g.Run()
```

Does this better than `sync.WaitGroup` ?

# Configuration

* [Configuration | Prometheus](https://prometheus.io/docs/prometheus/latest/configuration/configuration/)
* [prometheus/config - GoDoc](https://godoc.org/github.com/prometheus/prometheus/config)

```go
type Config struct {
    GlobalConfig   GlobalConfig    `yaml:"global"`
    AlertingConfig AlertingConfig  `yaml:"alerting,omitempty"`
    RuleFiles      []string        `yaml:"rule_files,omitempty"`
    ScrapeConfigs  []*ScrapeConfig `yaml:"scrape_configs,omitempty"`

    RemoteWriteConfigs []*RemoteWriteConfig `yaml:"remote_write,omitempty"`
    RemoteReadConfigs  []*RemoteReadConfig  `yaml:"remote_read,omitempty"`
    // contains filtered or unexported fields
}

func Load(s string) (*Config, error)
func LoadFile(filename string) (*Config, error)

func (c Config) String() string
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error
```

**Reload handler**




