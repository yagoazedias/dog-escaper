package infraestruture

import "fmt"

// Postgres  defines the parameters needed to connect to a Postgres database
type Postgres struct {
	Host     string `envconfig:"default=localhost"`
	Port     int    `envconfig:"default=5432"`
	Username string `envconfig:"default=postgres"`
	Password string `envconfig:"default=postgres"`
	Name     string `envconfig:"default=main"`

	MaxOpenConns int  `envconfig:"default=20"`
	LogMode      bool `envconfig:"default=false"`
}

// ConnectionString returns the connection string for a postgres database
func (p Postgres) ConnectionString() string {

	env := Postgres{Host: "localhost",
		Port: 5432,
		Username: "postgres",
		Password: "postgres",
		Name: "main",
		MaxOpenConns: 20,
		LogMode: false,
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.Username, env.Password, env.Name,
	)
}
