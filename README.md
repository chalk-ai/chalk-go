
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

1. From overrides configured by passing in a `*chalk.ClientConfig`
2. From environment variables if no arguments are passed 
3. From a ~/.chalk.yml file if neither 1 nor 2 are available

(2) and (3) are applicable only for these [options](https://github.com/chalk-ai/chalk-go/blob/main/internal/constants.go).
#### Without overrides
```go
import (
    "github.com/chalk-ai/chalk-go"
)

client := chalk.NewClient(context.Background())
```

#### With overrides
```go
client, err := chalk.NewClient(
	context.Background(),
	&chalk.ClientConfig{
		ClientId:      "id-89140a6614886982a6782106759e30",
		ClientSecret:  "sec-b1ba98e658d7ada4ff4c7464fb0fcee65fe2cbd86b3dd34141e16f6314267b7b",
		ApiServer:     "https://api.chalk.ai",
		EnvironmentId: "qa",
		Branch:        "jorges-december",
	}
)
```

### gRPC Client
To use gRPC as the underlying protocol for communication with Chalk: 
```go
// Create a client
client, err := chalk.NewGRPCClient(
	context.Background(),
	&chalk.GRPCClientConfig{Branch: "my-branch"},
)

// Online query
var users []User
res, err := chalk.OnlineQueryBulk(
	context.Background(),
	chalk.OnlineQueryParams{}.
		WithInput(Features.User.Id, []string{"u273489056"}).
		WithInput(Features.User.Transactions, [][]Transaction{
			{
				{Id: utils.ToPtr("txn8f76"), Amount: utils.ToPtr(13.23)},
				{Id: utils.ToPtr("txn546d"), Amount: utils.ToPtr(48.95)},
			},
		}).
		WithOutputs(Features.User.Id, Features.User.WeightedScore),
)
if err != nil {
	return errors.Wrap(err, "querying weighted score")
}
if err = res.UnmarshalInto(&users); err != nil {
	return errors.Wrap(err, "unmarshalling into users")
}
fmt.Println("user %s has weighted score %v", users[0].Id, users[0].WeightedScore)


// Multi-namespace online query
type underwriting struct {
	User 
	Loan  
}

res, err := chalk.OnlineQueryBulk(
    context.Background(),
    chalk.OnlineQueryParams{}.
        WithInput(Features.User.Id, []string{"u273489056"}).
        WithInput(Features.Loan.Id, []string{"l273489056"}).
        WithOutputs(
            Features.User.Id, 
            Features.User.WeightedScore, 
            Features.Loan.Id, 
            Features.Loan.ApprovalStatus,
        ),
)
if err != nil {
	return errors.Wrap(err, "querying weighted score and loan approval status")
}

var root []underwriting
if err = res.UnmarshalInto(&root); err != nil {
	return errors.Wrap(err, "unmarshalling into underwriting")
}
fmt.Println("user %s has weighted score %v", root[0].User.Id, root[0].User.WeightedScore)
fmt.Println("loan %s has approval status %v", root[0].Loan.Id, root[0].Loan.ApprovalStatus)
```

### Online Query

Query online features using the generated feature structs.  Access the results in the returned object or by passing the address of a variable with the correct type.

```go
user := User{}
_, err = client.OnlineQuery(
    context.Background(),
    chalk.OnlineQueryParams{}.
        WithInput(Features.User.Id, "u273489057").
        WithInput(Features.User.Transactions, []Transaction{
            {Id: utils.ToPtr("sd8f76"), Amount: utils.ToPtr(13.23)},
            {Id: utils.ToPtr("jk546d"), SeriesId: utils.ToPtr(48.95)},
        }).
        WithOutputs(Features.User.Id, Features.User.LastName),
    &user,
)
```

### Named Queries

If your deployment contains [named queries](https://docs.chalk.ai/docs/best-practices#create-named-queries-for-your-commonly-executed-queries),
you can specify a query name instead of outputs when making a query. 

```go
user := User{}
_, err = client.OnlineQuery(
    context.Background(),
    chalk.OnlineQueryParams{}.
        WithInput(Features.User.Id, "u273489057").
        WithQueryName("user_underwriting_features"),
    &user,
)
```



### Offline Query

When executing an offline query, a dataset is returned and can be downloaded as parquet files using the `DownloadData` method.

```go
res, _ := client.OfflineQuery(
    context.Background(),
    chalk.OfflineQueryParams{}.
        WithInput(Features.User.Id, []any{...}).
        WithOutputs(Features.User),
)

err = res.Revisions[0].DownloadData(<FILE_DIRECTORY>)
```


### Upload Features

Chalk allows you to synchronously persist features directly to your online and offline stores.

```go
res, err := client.UploadFeatures(
    context.Background(), 
    chalk.UploadFeaturesParams{
        Inputs: map[any]any{
            Features.User.Id: []string{"user-1", "user-2"},
            "user.last_name": []string{"Borges", "Paris"},
        },
    },
)
```

### Update Aggregates

Note: This method is available only when using the gRPC client. See [gRPC Client](#grpc-client) for more information.

You can easily update your windowed aggregation feature values with `UpdateAggregates`. For example, with this feature
definition:
```python
@features
class User:
    id: str
    txns: "DataFrame[Transaction]"
    txn_amount_total: Windowed[int] = windowed(
        "30d",
        "90d",
        materialization={
            "bucket_duration": "1d",
        },
        expression=_.txns[_.amount].sum(),
    )

@features
class Transaction:
    id: Primary[str]
    amount: float
    user_id: str
```
Then to update the `txn_amount_total` feature, you would upload features corresponding to that aggregation:
```go
res, err := client.UpdateAggregates(
    context.Background(),
    chalk.UpdateAggregatesParams{
        Inputs: map[any]any{
            "transaction.id": []string{"txn-1", "txn-2"},
            "transaction.user_id": []string{"user-1", "user-2"},
            "transaction.amount": []float64{100.0, 200.0},
            "transaction.__chalk_observed_at__": []time.Time{time.Now(), time.Now()},
        },
    },
)
```
Note that if you have an explicit `FeatureTime` feature specified, you could provide that in place of the 
`__chalk_observed_at__` column.  

### Querying against a branch

To query against a branch, create a `ChalkClient` with a `Branch` specified, and then make queries using that client.
```go
client, err := chalk.NewClient(
	context.Background(),
	&chalk.ClientConfig{
		Branch:        "jorges-december",
	},
)
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
