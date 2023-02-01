package postgres

type Option func(c *config)

func WithHost(host string) Option {
	return func(c *config) {
		c.host = host
	}
}

func WithPort(port int) Option {
	return func(c *config) {
		c.port = port
	}
}

func WithCreds(username, password string) Option {
	return func(c *config) {
		c.user = username
		c.password = password
	}
}

func WithDatabaseName(database string) Option {
	return func(c *config) {
		c.databaseName = database
	}
}

func WithSSLMode(enable bool) Option {
	return func(c *config) {
		c.sslMode = enable
	}
}
