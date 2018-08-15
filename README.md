API GO lucky
====

![alt text](https://github.com/sbasile-ch/api-go-lucky/blob/master/doc/images/img1.png "First working state")

Application to interact with the Companies House APIs.
- [CH API](https://developer.companieshouse.gov.uk/api/docs/index.html).

It allows to register for alerts anytime a change in the data of the registerird-for Company occurs.

Requirements
------------

- [Golang](https://golang.org/doc/install).

Getting Started
---------------

1. Make sure your `GOPATH` environment variable is set and its `bin` subdirectory
is included in `PATH` e.g.:
```shell
export GOPATH=${GOPATH:-$HOME/go}
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

2. Install:
```shell
go get github.com/sbasile-ch/api-go-lucky
api-go-lucky
```

## Configuration

The essential configuration item is:

```bash
export MY_CH_API=<your CH API key>
```

## TODO
- Add alert trigger and logging/display Panel
- Add History Panel
- Add tests
- Add validation
- Add Authentication
- ...

