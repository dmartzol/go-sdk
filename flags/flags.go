package flags

import "github.com/urfave/cli"

var (
	HostnameFlagName           = "host"
	PortFlagName               = "port"
	LogsFormatFlagName         = "logsFormat"
	RawRequestsLoggingFlagName = "logRawRequests"
	DatabaseHostnameFlagName   = "databaseHostname"
	DatabasePortFlagName       = "databasePort"
	DatabaseUserFlagName       = "databaseUser"
	DatabaseNameFlagName       = "databaseName"
	DatabasePasswordFlagName   = "databasePassword"

	CommonFlags = []cli.Flag{
		&cli.StringFlag{
			Name:   LogsFormatFlagName,
			EnvVar: "LOGS_FORMAT",
			Value:  "console",
		},
		&cli.StringFlag{
			Name:   DatabaseHostnameFlagName,
			EnvVar: "PGHOST",
			Value:  "localhost",
		},
		&cli.StringFlag{
			Name:   DatabasePortFlagName,
			EnvVar: "PGPORT",
			Value:  "5432",
		},
		&cli.StringFlag{
			Name:   DatabaseUserFlagName,
			EnvVar: "PGUSER",
			Value:  "",
		},
		&cli.StringFlag{
			Name:   DatabaseNameFlagName,
			EnvVar: "PGNAME",
			Value:  "",
		},
		&cli.StringFlag{
			Name:   DatabasePasswordFlagName,
			EnvVar: "PGPASSWORD",
			Value:  "password",
		},
	}
)
