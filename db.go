package config

import "fmt"

type Postgres struct {
	Host string `config:"HOST,default=localhost"`
	Port string `config:"PORT,default=5432"`

	User string `config:"USER,default=postgres"`
	Pass string `config:"PASS,default=postgres"`
	Name string `config:"NAME,default=postgres"`

	SSLMode string `config:"SSL_MODE,default=disable"`

	DSN string `config:"DSN"`
}

func (c *Postgres) GetDSN() string {
	if c.DSN != "" {
		return c.DSN
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Pass, c.Host, c.Port, c.Name, c.SSLMode)
}
