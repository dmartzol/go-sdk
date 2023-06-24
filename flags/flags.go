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

	// tracing flags
	TracingEnabledFlag    = "tracingEnabled"
	OtelCollectorHostFlag = "tracingHost"
	OtelCollectorPortFlag = "tracingPort"

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
		&cli.BoolFlag{
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

	TracingFlags = []cli.Flag{
		&cli.BoolFlag{
			Name:   TracingEnabledFlag,
			EnvVar: "TRACING_ENABLED",
		},
		&cli.StringFlag{
			Name:   OtelCollectorHostFlag,
			EnvVar: "OTEL_COLLECTOR_HOST",
			Value:  "otel",
		},
		&cli.StringFlag{
			Name:   OtelCollectorPortFlag,
			EnvVar: "OTEL_COLLECTOR_PORT",
			Value:  "4317",
		},
	}
)
