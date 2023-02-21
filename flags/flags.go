package flags

import "github.com/urfave/cli"

var (
	// Server flags
	HostnameFlagName = "host"
	PortFlagName     = "port"

	// Logger flags
	LogsFormatFlagName         = "logsFormat"
	RawRequestsLoggingFlagName = "logRawRequests"

	// Database flags
	DatabaseHostnameFlag = "databaseHostname"
	DatabasePortFlag     = "databasePort"
	DatabaseNameFlag     = "databaseName"
	DatabaseUserFlag     = "databaseUser"
	DatabasePasswordFlag = "databasePassword"

	// scraper flags
	LocalScrapeFlag = "localScrape"
	seleniumHubPort = "seleniumHubPort"
	seleniumHubHost = "seleniumHubHost"

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
			Name:   RawRequestsLoggingFlagName,
			Usage:  "Log incoming raw http requests",
			EnvVar: "LOG_RAW_REQUESTS",
		},
		&cli.StringFlag{
			Name:   LogsFormatFlagName,
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
		&cli.BoolTFlag{
			Name:   LocalScrapeFlag,
			EnvVar: "LOCAL_SCRAPE",
		},
		&cli.IntFlag{
			Name:   seleniumHubPort,
			EnvVar: "SELENIUM_HUB_PORT",
			Value:  4444,
		},
		&cli.StringFlag{
			Name:   seleniumHubHost,
			EnvVar: "SELENIUM_HUB_HOST",
			Value:  "selenium-hub",
		},
	}
)
