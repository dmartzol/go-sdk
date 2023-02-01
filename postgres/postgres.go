package postgres

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type config struct {
	host         string
	port         int
	user         string
	password     string
	databaseName string
	serviceName  string
}

func (c *config) url() string {
	query := url.Values{}

	query.Set("client_encoding", "UTF8")

	if c.serviceName != "" {
		query.Set("application_name", c.serviceName)
	}

	datasource := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.user, c.password),
		Host:     net.JoinHostPort(c.host, strconv.Itoa(c.port)),
		Path:     c.databaseName,
		RawQuery: query.Encode(),
	}

	return datasource.String()
}

func New(serviceName string, options ...Option) (*sqlx.DB, error) {
	// default config options go here
	c := &config{
		host:        "localhost",
		serviceName: serviceName,
		port:        5432,
	}

	for _, opt := range options {
		opt(c)
	}

	db, err := sqlx.Connect("postgres", c.url())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
