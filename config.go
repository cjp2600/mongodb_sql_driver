package mongodb_sql_driver

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	MongoDSN    string
	Timeout     time.Duration
	PingTimeout time.Duration
	Location    *time.Location
	Debug       bool
	Params      map[string]string
}

// NewConfig creates a new config with default values
func NewConfig() *Config {
	return &Config{
		Timeout:     10 * time.Second,
		PingTimeout: 2 * time.Second,
		Location:    time.UTC,
		Params:      make(map[string]string),
	}
}

func Parse(dsn string) (*Config, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	cfg := NewConfig()

	if err = parsePrams(cfg, map[string][]string(u.Query())); err != nil {
		return nil, err
	}
	return cfg, nil
}

func parsePrams(cfg *Config, params map[string][]string) (err error) {
	for k, v := range params {
		if len(v) == 0 {
			continue
		}
		switch strings.ToLower(k) {
		case "timeout":
			cfg.Timeout, err = time.ParseDuration(v[0])
		case "pingtimeout":
			cfg.PingTimeout, err = time.ParseDuration(v[0])
		case "location":
			cfg.Location, err = time.LoadLocation(v[0])
		case "debug":
			cfg.Debug, err = strconv.ParseBool(v[0])
		case "mongodsn":
			cfg.MongoDSN = v[0]
		default:
			cfg.Params[k] = v[0]
		}
		if err != nil {
			return err
		}
	}

	return
}
