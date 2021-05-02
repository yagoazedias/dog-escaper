package infraestruture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}


// Postgres  defines the parameters needed to connect to a Postgres database
type Postgres struct {
	Host     string `envconfig:"default=localhost"`
	Port     string    `envconfig:"default=5432"`
	Username string `envconfig:"default=postgres"`
	Password string `envconfig:"default=postgres"`
	Name     string `envconfig:"default=main"`

	MaxOpenConns int  `envconfig:"default=20"`
	LogMode      bool `envconfig:"default=false"`
}

// ConnectionString returns the connection string for a postgres database
func (p Postgres) ConnectionString() string {

	env := Postgres{
		Host: Config["PG_HOST"],
		Port: Config["PG_PORT"],
		Username: Config["PG_DB_USERNAME"],
		Password: Config["PG_PASSWORD"],
		Name: Config["PG_DB_NAME"],
		MaxOpenConns: 20,
		LogMode: false,
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.Host, env.Port, env.Username, env.Password, env.Name,
	)
}

var Config = map[string]string{
	"MQTT_HOST": os.Getenv("MQTT_HOST"),
	"MQTT_PORT": os.Getenv("MQTT_PORT"),
	"MQTT_TOPIC": os.Getenv("MQTT_TOPIC"),
	"PG_HOST": os.Getenv("PG_HOST"),
	"PG_DB_NAME": os.Getenv("PG_DB_NAME"),
	"PG_DB_USERNAME": os.Getenv("PG_DB_USERNAME"),
	"PG_PASSWORD": os.Getenv("PG_PASSWORD"),
	"PG_PORT": os.Getenv("PG_PORT"),
}
