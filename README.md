# go-sdk

go-sdk is a personal Golang SDK that provides common functionality for personal projects, such as loading configuration values, database connections, and logging.

## Installation

To use go-sdk in your Golang project, you can install it using `go get`:

```
go get github.com/dmartzol/go-sdk
```

## Usage

```go
import (
	"fmt"
	"net/http"
	"os"

	"github.com/dmartzol/go-sdk/flags"
	"github.com/dmartzol/go-sdk/logger"
	"github.com/dmartzol/go-sdk/postgres"
	"github.com/dmartzol/hmm/cmd/backend/api"
	"github.com/urfave/cli"
)

const (
	appName = "backend"
)

func main() {
	app := &cli.App{
		Name:   appName,
		Action: newBackendServiceRun,
	}
	app.Flags = append(app.Flags, flags.DatabaseFlags...)
	app.Flags = append(app.Flags, flags.LoggerFlags...)
	app.Flags = append(app.Flags, flags.ServerFlags...)

	err := app.Run(os.Args)
	if err != nil {
		sdkLogger := logger.New()
		sdkLogger.Errorf("error running app: %v", err)
		os.Exit(1)
	}

}

func newBackendServiceRun(c *cli.Context) error {
	postgresOpts := []postgres.Option{
		postgres.WithHost(c.String(flags.DatabaseHostnameFlag)),
		postgres.WithDatabaseName(c.String(flags.DatabaseNameFlag)),
		postgres.WithCreds(
			c.String(flags.DatabaseUserFlag),
			c.String(flags.DatabasePasswordFlag),
		),
	}
	db, err := postgres.New(appName, postgresOpts...)
	if err != nil {
		return fmt.Errorf("unable to initialize database: %w", err)
	}
	defer db.Close()

	// Initializes a new logger using provided configuration and options.
	loggerOpts := []logger.Option{
		logger.WithLevel(logger.Level(c.String(flags.LogsLevelFlag))),
		logger.WithEncoding(logger.Encoding(c.String(flags.LogsFormatFlag))),
	}
	sdkLogger := logger.NewWithOptions(loggerOpts...)

	restAPI := api.NewAPI(db, sdkLogger)

	host := c.String(flags.HostnameFlagName)
	port := c.String(flags.PortFlagName)
	address := host + ":" + port
	server := &http.Server{
		Addr:    address,
		Handler: restAPI,
	}

	restAPI.Logger.Info("starting server", "address", address)
	return server.ListenAndServe()
}
```

