# wygop

*wygop* is a Go client library for accessing the [Wykop API v2][].

## Installation ##

*wygop* is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/cymruu/wygop
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/cymruu/wygop"
```

and run `go get` without parameters.

## Usage ##

```go
import "github.com/cymruu/wygop"
```

Construct a new *wygop* client, then use the various services on the client to
access different parts of the Wykop API. For example:


```go
	flag.Parse()
	ctx := context.Background()

	client := wygop.CreateClient(appkey, secret, http.DefaultClient)
	service := wygops_service.CreateWykopService(client)

	entry, err := service.Entries.Entry(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", entry.Body)
```

NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.

For more sample code snippets, head over to the
[example](https://github.com/cymruu/wygop/tree/master/example) directory.

### Authentication ###

To access endpoints which require authentication, special "connection" key is required.
The key can be obtained from this [site](https://www.wykop.pl/dla-programistow/apiv2/).


```go
func main() {
	ctx := context.Background()

	client := wygop.CreateClient(appkey, secret, http.DefaultClient)
	service := wygops_service.CreateWykopService(client)

	_, err := notifierService.Login.Index(ctx, connectionKey)
	if err != nil {
		log.Fatal(err)
	}
	//client is now authenticated and can access protected endpoints
}
```


[Wykop API v2]: https://www.wykop.pl/dla-programistow/apiv2docs/

## Contributing ##
I would like to cover contributions are of course always welcome. Just create new issue or open a Pull Request.

## Versioning ##

In general, *wygop* follows [semver](https://semver.org/).
## License ##

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
