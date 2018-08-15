API GO lucky
====

![alt text](https://github.com/sbasile-ch/api-go-lucky/tree/master/doc/images/img1.png "First working state")

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

2. Git clone gotools and eric:
```shell
mkdir -p $GOPATH/src/github.com/sbasile-ch/api-go-lucky
cd !$
git clone git@github.com/sbasile-ch/api-go-lucky.git
cd api-go-lucky
go build && ./api-go-lucky
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

