# EdgeX Foundry REST Application service

## Table of content

- [EdgeX Foundry REST Application service](#edgex-foundry-rest-application-service)
  - [Table of content](#table-of-content)
  - [About this project](#about-this-project)
  - [Getting started](#getting-started)
    - [adding to EdgeX Foundry docker-compose (recommended)](#adding-to-edgex-foundry-docker-compose-recommended)
    - [docker-compose](#docker-compose)
    - [binary](#binary)
    - [source code](#source-code)
  - [Config](#config)

## About this project

This is an application service for the [EdgeX Foundry platform](https://www.edgexfoundry.org/). It provides a RESTful API to get information about devices, device readings and to use device commands.
It also provides access control based on users, roles and policies.

This project was developed as part of my bachelor thesis in cooperation with Faculty 07 of the TH Köln.

## Getting started

Make sure the [EdgeX Foundry platform](https://www.edgexfoundry.org/) is running. This service is tested with 1.2-Geneva.

### adding to EdgeX Foundry docker-compose (recommended)

Add the following to the services:

```yaml
  edgex-restapp:
    image: mmomper/edgex-restapp
    container_name: edgex-app-service-rest
    hostname: edgex-app-service-rest
    ports:
      - 8080:8080
      - 8443:8443
    networks:
      - edgex-network
    volumes:
      - persist:/persist
    environment:
      ADMIN_USERNAME: admin
      ADMIN_PASSWORD: "strongpassword"
      CLIENT_COREDATA_HOST: edgex-core-data
      CLIENT_COREMETADATA_HOST: edgex-core-metadata
      CLIENT_CORECOMMAND_HOST: edgex-core-command
    # command: /edgex-restapp --scheme http
```

Also add this to volumes:

```yaml
persist:
```

Now start the compose as usual.

As default, it starts listening http and https. To only use one, uncomment the following line. Change http to https if you want https.

```yaml
command: /edgex-restapp --scheme http
```

### docker-compose

Make sure you have docker and docker-compose (included with docker desktop) installed.

Download the docker-compose.yml from [here](https://raw.githubusercontent.com/Momper14/edgex-restapp/master/docker-compose.yml).

Set the network name from the EdgeX Foundry platform docker network. Remove the entry if it doesn't run on the local docker.

```yaml
networks:
  edgex-network:
    external:
      name: edgex-network
```

If you don't know the network name, look it up with

```sh
docker network ls
```

or set the name in the EdgeX Foundry docker-compose.yml (don't forget to restart it!)

```yaml
networks:
  edgex-network:
    driver: "bridge"
    name: edgex-network
```

Run the service with

```sh
docker-compose up -d
```

As default, it starts listening http and https. To only use one, add (uncomment)

```yaml
command: /edgex-restapp --scheme http
```

to the docker-compose.yml file. Change http to https if you want https.

### binary

Download the appropriate binary for your system and run them with a command shell.

As default, it starts listening http and https. To only use one, add the arg --scheme http or --scheme https

### source code

First, make sure the golang runtime with at least version 1.15 is installed.

Next, get the source code. For example with git:

```sh
git clone https://github.com/Momper14/edgex-restapp.git
cd edgex-restapp
```

After that, just run it with

```sh
go run ./cmd/edgex-restapp-server/
```

As default, it starts listening http and https. To only use one, add the arg --scheme http or --scheme https

## Config

This service can be configured with a yaml file or with environment variables. The default yaml file is config.yml. This can be changed with the environment variable CONFIG_FILE.

To use environment variables, the key must be uppercase and contain "_" (underscore) instead of "." (dots). For example: to set admin.password, use ADMIN_PASSWORD="strongpassword".

Environment variables overrides configurations from the yaml file. It means, the hierarchy is: environment variable -> config yaml -> default.

Configurations are as following:

| Key                          | Default             | Description                                                        |
| ---------------------------- | ------------------- | ------------------------------------------------------------------ |
| host                         | 0.0.0.0             | ip to listen for http requests.                                    |
| port                         | 8080                | port to listen for http requests.                                  |
| tls.host                     |                     | ip to listen for https requests. Uses "host" if not specified.     |
| tls.port                     | 8443                | port to listen for https requests.                                 |
| tls.certificate              | tls/certificate.crt | path to tls certificate (.crt).                                    |
| tls.key                      | tls/key.key         | path to tls key (.key).                                            |
| tls.ca.certificate           |                     | path to tls ca certificate.                                        |
| admin.username               | admin               | username for admin. Works only on first start (db does not exist). |
| admin.password               | password            | password for admin. Works only on first start (db does not exist). |
| admin.role                   | admin               | role for admin. Works only on first start (db does not exist).     |
| role.default                 | guest               | role which will be used for requests without auth.                 |
| client.coredata.protocol     | http                | protocol which the core data service uses to communicate.          |
| client.coredata.host         | localhost           | host of the core data service.                                     |
| client.coredata.port         | 48080               | port of the core data service.                                     |
| client.coremetadata.protocol | http                | protocol which the core metadata service uses to communicate.      |
| client.coremetadata.host     | localhost           | host of the core metadata service.                                 |
| client.coremetadata.port     | 48081               | port of the core metadata service.                                 |
| client.corecommand.protocol  | http                | protocol which the core command service uses to communicate.       |
| client.corecommand.host      | localhost           | host of the core command service.                                  |
| client.corecommand.port      | 48082               | port of the core command service.                                  |
| db.userdb                    | persist/user.db     | file to store the user db.                                         |
| db.roledb                    | persist/role.db     | file to store the role db.                                         |
| enforcer.model               | persist/model.conf  | file to use as model for the casbin enforcer.                      |
| enforcer.policy              | persist/policy.csv  | file to store the policies for the casbin enforcer.                |
