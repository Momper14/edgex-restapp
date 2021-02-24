# Readme for devs

## Table of content

- [Readme for devs](#readme-for-devs)
  - [Table of content](#table-of-content)
  - [introduction](#introduction)
  - [used lib's](#used-libs)
    - [skv](#skv)
    - [resty](#resty)
    - [logrus](#logrus)
    - [viper](#viper)
    - [negroni](#negroni)
    - [casbin](#casbin)
  - [package overview](#package-overview)
  - [notable packages](#notable-packages)
    - [client](#client)
    - [db](#db)
    - [restapi](#restapi)
  - [changing a model](#changing-a-model)

## introduction

This readme helps to get a quick overview of the service and its structure.

This service is generated with the [go-swagger](https://github.com/go-swagger/go-swagger) tool. This tool generates a server with all defined models, path's and responses from a swagger definition. Only the handlers itself must be defined. At the moment it does support only swagger 2.0.

## used lib's

### skv

[skv](https://github.com/Momper14/skv) is a simple key-value store, which uses a file to store the data. It is based on the [bbolt](https://github.com/etcd-io/bbolt) database.

### resty

[resty](https://github.com/go-resty/resty) is a simple HTTP and REST client library for Go.

### logrus

[logrus](https://github.com/sirupsen/logrus) is a structured logger for Go, completely API compatible with the standard library logger.

### viper

[viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

### negroni

[negroni](https://github.com/urfave/negroni) is an idiomatic approach to web middleware in Go. It is tiny, non-intrusive, and encourages use of net/http Handlers.

### casbin

[casbin](https://github.com/casbin/casbin) is a powerful and efficient open-source access control library for Golang projects. It provides support for enforcing authorization based on various access control models.

## package overview

```text
.                                   root folder
├── client                          client to communicate with the other services
│   └── models                      models of the client
├── cmd                             executables
│   └── edgex-restapp-server        server executable (generated)
├── db                              provides functions to communicate with the databases
│   └── models                      database models
├── init                            init functions to initialize things like logger etc.
├── models                          models of the api (generated)
├── openapi                         location of the swagger spec yaml
├── restapi                         package to configure the api server including authentication and authorization (partially generated)
│   ├── json                        json files for tests
│   ├── operations                  operations the api can do (generated)
│   ├── converter                   provides functions to convert between api models and client/db models.
└── util                            utilities like custom errors for the api server
```

## notable packages

### client

This package provides functions to communicate with other services. It uses resty clients to to so.

Is has also a custom error for response errors and some functions to help with validating and processing of the response.

### db

This package provides functions to communicate with the databases. Roles and users are stored with a skv db for each. Policies are managed over a casbin Enforcer and stored in a csv file.

### restapi

This package contains the server with it's configuration. doc.go, embedded_spec.go, server.go and everything under operations/ are generated and should not be edited.

Configuration should be done in the configure_*.go files. For example, configure_users.go contains a function to configure the handler for the /users path. configure_edgex_restapp.go contains the central configuration and calls the other configure functions.

authentication.go provides a function, which will registered on the generated server to handle authentication for a request.

authorization.go provides a function, which will registered on the generated server to handle authorization for a request. It uses the default role if no participant is given (no auth header was in the request).

## changing a model

To change a model, change it under the definitions in the swagger-file, run

```go
go generate ./...
```

and modify the corresponding converter function in the converter package.
