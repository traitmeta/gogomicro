# Inventory Service

This is the Inventory service

Generated with

```
micro new github.com/songxuexian/gogomicro/inventory-srv --namespace=sxx.micro.book --alias=inventory --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: sxx.micro.book.srv.inventory
- Type: srv
- Alias: inventory

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
./inventory-srv
```

Build a docker image
```
make docker
```