woothee-go
==========

[![Build Status](https://travis-ci.org/woothee/woothee-go.png?branch=master)](https://travis-ci.org/woothee/woothee-go)

[![Coverage Status](https://coveralls.io/repos/woothee/woothee-go/badge.png?branch=HEAD)](https://coveralls.io/r/woothee/woothee-go?branch=topic%2Fgoveralls)

[Woothee](https://github.com/woothee) is a multi-language UserAgent detection library, and this is the Go version of Woothee.

```go
import (
    "github.com/woothee/woothee-go"
)

func main() {
    agent := `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`
    result, err := woothee.Parse(agent)
    if err != nil {
        log.Fatalf("Could not parse '%s': %s", agent, err)
    }

    /*
        result.Name     = "Googlebot"
        result.Category = "crawler"
        result.Os       = "UNKNOWN"
        result.Type     = "UNKNOWN"
        result.Version  = "UNKNOWN"
        result.Vendor   = "UNKNOWN"
    */
}
```

# Working with the source code

Note: You need to have GOROOT, GOPATH et al set correctly

## Fetching the source

```
go get github.com/woothee/woothee-go
```

## Running Tests

```
go test -v ./...
```

## Updating woothee definition

If you live on the bleeding edge (i.e. go 1.4 and up):

```
go generate
```

Otherwise:

```
go get gopkg.in/yaml.v2
```

then

```
go run maint/gendataset.go
```

This updates the dataset.go and other tests
