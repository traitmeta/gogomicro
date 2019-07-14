# User Service

This is the User service

Generated with

```
micro new github.com/gogomicro/user-web --namespace=sxx.micro.book --alias=user --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: sxx.micro.book.web.user
- Type: web
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-web
```

Build a docker image
```
make docker
```