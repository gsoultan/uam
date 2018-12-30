package config

type Database struct {
	Dialect  string `env:"USER_DATABASE_DIALECT"`
	Host     string `env:"USER_DATABASE_HOST"`
	Port     string `env:"USER_DATABASE_PORT"`
	Name     string `env:"USER_DATABASE_NAME"`
	User     string `env:"USER_DATABASE_USERNAME"`
	Password string `env:"USER_DATABASE_PASSWORD"`
}
