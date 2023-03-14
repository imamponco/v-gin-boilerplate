package contract

import (
	"net/url"
)

type Config struct {
	Debug                   string  `envconfig:"DEBUG"`
	Port                    int     `envconfig:"PORT"`
	Stage                   string  `envconfig:"STAGE"`
	ServerSecure            bool    `envconfig:"SERVER_LISTEN_SECURE"`
	ServerCert              string  `envconfig:"SERVER_CERT_PATH"`
	ServerCertKey           string  `envconfig:"SERVER_CERT_PATH_KEY"`
	ServerTrustProxy        string  `envconfig:"SERVER_TRUST_PROXY"`
	ServerBaseURL           url.URL `envconfig:"SERVER_HTTP_BASE_URL"`
	CORS                    string  `envconfig:"CORS"`
	DatabaseDriver          string  `envconfig:"DB_DRIVER"`
	DatabaseHost            string  `envconfig:"DB_HOST"`
	DatabasePort            string  `envconfig:"DB_PORT"`
	DatabaseUsername        string  `envconfig:"DB_USER"`
	DatabasePassword        string  `envconfig:"DB_PASS"`
	DatabaseName            string  `envconfig:"DB_NAME"`
	DatabaseMaxIdleConn     string  `envconfig:"DB_POOL_MAX_IDLE_CONN"`
	DatabaseMaxOpenConn     string  `envconfig:"DB_POOL_MAX_OPEN_CONN"`
	DatabaseMaxConnLifetime string  `envconfig:"DB_POOL_MAX_CONN_LIFETIME"`
	DatabaseBootMigration   string  `envconfig:"DB_BOOT_MIGRATION"`
	JWTKey                  string  `envconfig:"JWT_KEY"`
	JWTExpiry               int     `envconfig:"JWT_EXPIRY"`
}

func (c *Config) GetHTTPBaseURL() string {
	u := c.ServerBaseURL
	return u.String()
}
