
[![Go Reference](https://pkg.go.dev/badge/github.com/chalk-ai/chalk-go.svg)](https://pkg.go.dev/github.com/chalk-ai/chalk-go)
[![Test](https://github.com/chalk-ai/chalk-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/chalk-ai/chalk-go/actions/workflows/test.yml)

# Go Chalk

The official [Chalk](https://chalk.ai) client library.

## Usage
### Requirements

- Go 1.19 or later

### Installation

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
### Codegen

Chalk generates Go structs from your python feature definitions, which makes it easy to use your features from Go.

Run the code-gen command inside your chalk project to generate a file containing your generated structs, then copy that file into your go project.

```sh
chalk codegen go --out=<OUTPUT_FILEPATH> 
```


### Connect to chalk

Create a client using the `NewClient` method.  The returned client gets its configuration:

1. From the contructor arguments
2. From environment variables if no arguments are passed
3. From a ~/.chalk.yml file if neither 1 nor 2 are available
```go
import (
    "github.com/chalk-ai/chalk-go"
)

client := chalk.NewClient()
```

### Online Query

Query online features using the generated feature structs.  Access the results in the returned object or by passing the address of a variable with the correct type.

```go
user := User{}
_, err = client.OnlineQuery(
    chalk.OnlineQueryParams{}.
        WithInput(Features.User.Id, "<INPUT_VALUE>").
        WithOutputs(Features.User.Id, Features.User.LastName),
    &user,
)
```


### Offline Query

When executing an offline query, a dataset is returned and can be downloaded as parquet files using the `DownloadData` method.

```py
res, _ := client.OfflineQuery(
    chalk.OfflineQueryParams{}.
        WithInput(Features.User.Id, []any{...}).
        WithOutputs(Features.User),
)

err = res.Revisions[0].DownloadData(<FILE_DIRECTORY>)
```

### Configuring Logging

By default, Chalk logs error messages only (which are sent to `stderr`).
Configure default logging using the global `DefaultLeveledLogger` variable:

```go
chalk.DefaultLeveledLogger = &chalk.StdOutLeveledLogger{
    Level: chalk.LevelInfo,
}
```

<!-- Or on a per-client basis:

```go
config := &chalk.ClientConfig{
    Logger: &chalk.StdOutLeveledLogger{
        Level: chalk.LevelInfo,
    },
}
client := chalk.Client(config)
``` -->

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


## Contributing

We'd love to accept your patches!  If you submit a pull request, please keep the following guidelines in mind:

1. Fork the repo, develop and test your code changes.
2. Ensure your code adheres to the existing style. 
3. Ensure that your code has an appropriate set of tests that pass.
4. Submit a pull request.


### Development

Clone the git repo from github.com/chalk-ai/chalk-go.

1. All patches must be `go fmt` compatible.
2. All types, structs, and functions should be documented.

### Testing

To execute the tests: `go test ./...`.  Enure all tests pass

## License

Apache 2.0 - See [LICENSE](LICENSE) for more information.
