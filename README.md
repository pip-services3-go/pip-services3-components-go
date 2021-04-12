# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> Component definitions for Pip.Services in Gol

This module is a part of the [Pip.Services](http://pipservices.org) polyglot microservices toolkit.

The Components module contains standard component definitions that can be used to build applications and services.

The module contains the following packages:

- [**Auth**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/auth) - authentication credential stores
- [**Build**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/build) - factories
- [**Cache**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/cache) - distributed cache
- [**Component**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/component) - the root package
- [**Config**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/config) - configuration readers
- [**Connect**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/connect) - connection discovery services
- [**Count**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/count) - performance counters
- [**Info**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/info) - context info
- [**Lock**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/lock) - distributed locks
- [**Log**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/log) - logging components
- [**Test**](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/test) - test components

<a name="links"></a> Quick links:

* [Logging](https://www.pipservices.org/recipies/logging)
* [Configuration](https://www.pipservices.org/recipies/configuration) 
* [API Reference](https://godoc.org/github.com/pip-services3-go/pip-services3-components-go/)
* [Change Log](CHANGELOG.md)
* [Get Help](https://www.pipservices.org/community/help)
* [Contribute](https://www.pipservices.org/community/contribute)


## Use

Get the package from the Github repository:
```bash
go get -u github.com/pip-services3-go/pip-services3-components-go@latest
```

## Develop

For development you shall install the following prerequisites:
* Golang v1.12+
* Visual Studio Code or another IDE of your choice
* Docker
* Git

Run automated tests:
```bash
go test -v ./test/...
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized test as:
```bash
./test.ps1
./clear.ps1
```

## Contacts

The library is created and maintained by **Sergey Seroukhov**.

The documentation is written by **Levichev Dmitry**.
