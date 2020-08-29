package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Unmarshaler interface {
	UnmarshalConfig(b []byte) error
}

type Config struct {
	Logging LoggingConfig  `json:"logging"`
	Servers []ServerConfig `json:"servers"`
	Service ServiceConfig  `json:"service"`
}

type LoggingConfig struct {
	Debug    bool   `json:"debug"`
	Location string `json:"location"`
}

type ServerConfig struct {
	Kind   string      `json:"kind"`
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
}

func (sc *ServerConfig) UnmarshalJSON(b []byte) error {
	var j map[string]*json.RawMessage
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	var name string
	if err := json.Unmarshal(*j["name"], &name); err != nil {
		return err
	}

	sc.Name = name

	var kind string
	if err := json.Unmarshal(*j["kind"], &kind); err != nil {
		return err
	}

	sc.Kind = kind

	switch sc.Kind {
	case "irc":
		var irc IRCServerConfig

		if err := json.Unmarshal(*j["config"], &irc); err != nil {
			return err
		}

		irc.Name = name

		sc.Config = irc
	default:
		return fmt.Errorf("no such server kind %s", sc.Kind)
	}

	return nil
}

type IRCServerConfig struct {
	Name       string   `json:"-"`
	ServerAddr string   `json:"server_addr"`
	Password   string   `json:"password"`
	Nicks      []string `json:"nicks"`
	User       string   `json:"user"`
	RealName   string   `json:"real_name"`
	Channels   []string `json:"channels"`
	Commands   []string `json:"commands"`
	UseTLS     bool     `json:"use_tls"`
	RootCAPath string   `json:"root_ca"`
}

type ServiceConfig struct {
	BindAddr string
}

func NewConfigFromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
