package config

// Log defines the available logging configuration.
type Log struct {
	Level  string
	Pretty bool
	Color  bool
}

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string
	Token  string
	Pprof  bool
	Zpages bool
}

// HTTP defines the available http configuration.
type HTTP struct {
	Addr string
	Root string
}

// Tracing defines the available tracing configuration.
type Tracing struct {
	Enabled   bool
	Type      string
	Endpoint  string
	Collector string
	Service   string
}

// Ldap defined the avialable LDAP configuration.
type Ldap struct {
	Network      string
	Address      string
	UserName     string
	Password     string
	BaseDNUsers  string
	BaseDNGroups string
}

// Config combines all available configuration parts.
type Config struct {
	File    string
	Log     Log
	Debug   Debug
	HTTP    HTTP
	Tracing Tracing
	Ldap    Ldap
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}
