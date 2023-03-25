# Go Chalk

The official [Chalk](https://chalk.ai) client library.

## Requirements

- Go 1.19 or later

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference `chalk-go` in a Go program with `import`:

``` go
import (
    "github.com/chalk-ai/chalk-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the `chalk-go` module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get -u github.com/chalk-ai/chalk-go
```

## Quickstart

```go
import (
    "github.com/chalk-ai/chalk-go"
)

client := chalk.NewClient()
```



## Configuring Logging

By default, Chalk logs error messages only (which are sent to `stderr`).
Configure default logging using the global `DefaultLeveledLogger` variable:

```go
chalk.DefaultLeveledLogger = &chalk.StdOutLeveledLogger{
    Level: chalk.LevelInfo,
}
```

Or on a per-client basis:

```go
config := &chalk.ClientConfig{
    Logger: &chalk.StdOutLeveledLogger{
        Level: chalk.LevelInfo,
    },
}
client := chalk.Client(config)
```

It's possible to use non-Chalk leveled loggers as well. Chalk expects loggers
to comply to the following interface:

```go
type LeveledLogger interface {
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
}
```

Some loggers like Logrus and Zap's `SugaredLogger`
support this interface natively, so it's possible to set
`DefaultLeveledLogger` to a `*logrus.Logger` or `*zap.SugaredLogger` directly.
To use other loggers, you may need a shim layer.
