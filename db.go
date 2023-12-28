package config

import "fmt"

type Postgres struct {
	Host string `config:"DB_HOST,default=localhost,required"`
	Port string `config:"DB_PORT,default=5432,required"`

	User string `config:"DB_USER,default=postgres,required"`
	Pass string `config:"DB_PASS,default=postgres,required"`
	Name string `config:"DB_NAME,default=postgres,required"`

	SSLMode string `config:"DB_SSL_MODE,default=disable,required"`

	DSN string `config:"DSN"`
}

func (c *Postgres) GetDSN() string {
	if c.DSN != "" {
		return c.DSN
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Pass, c.Host, c.Port, c.Name, c.SSLMode)
}
