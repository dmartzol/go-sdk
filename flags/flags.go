package flags

import "github.com/urfave/cli"

var (
	// Server flags
	HostnameFlagName = "host"
	PortFlagName     = "port"

	// Logger flags
	LogsFormatFlag         = "logsFormat"
	RawRequestsLoggingFlag = "logRawRequests"

	// Database flags
	DatabaseHostnameFlag = "databaseHostname"
	DatabasePortFlag     = "databasePort"
	DatabaseNameFlag     = "databaseName"
	DatabaseUserFlag     = "databaseUser"
	DatabasePasswordFlag = "databasePassword"

	// scraper flags
	SeleniumHubPortFlag = "seleniumHubPort"
	SeleniumHubHostFlag = "seleniumHubHost"

	ServerFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   HostnameFlagName,
			EnvVar: "HOSTNAME",
			Value:  "localhost",
		},
		&cli.IntFlag{
			Name:   PortFlagName,
			EnvVar: "PORT",
			Value:  8080,
		},
	}

	LoggerFlags = []cli.Flag{
		&cli.BoolTFlag{
			Name:   RawRequestsLoggingFlag,
			Usage:  "Log incoming raw http requests",
			EnvVar: "LOG_RAW_REQUESTS",
		},
		&cli.StringFlag{
			Name:   LogsFormatFlag,
			EnvVar: "LOGS_FORMAT",
			Value:  "console",
		},
	}

	DatabaseFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   DatabaseHostnameFlag,
			EnvVar: "PGHOST",
			Value:  "localhost",
		},
		&cli.StringFlag{
			Name:   DatabasePortFlag,
			EnvVar: "PGPORT",
			Value:  "5432",
		},
		&cli.StringFlag{
			Name:   DatabaseNameFlag,
			EnvVar: "PGNAME",
			Value:  "",
		},
		&cli.StringFlag{
			Name:   DatabaseUserFlag,
			EnvVar: "PGUSER",
			Value:  "",
		},
		&cli.StringFlag{
			Name:   DatabasePasswordFlag,
			EnvVar: "PGPASSWORD",
			Value:  "password",
		},
	}

	ScraperFlags = []cli.Flag{
		&cli.IntFlag{
			Name:   SeleniumHubPortFlag,
			EnvVar: "SELENIUM_HUB_PORT",
			Value:  4444,
		},
		&cli.StringFlag{
			Name:   SeleniumHubHostFlag,
			EnvVar: "SELENIUM_HUB_HOST",
			Value:  "selenium-hub",
		},
	}
)
