package config

import "net"

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func (c RedisConfig) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
}
